/*
 * Copyright (c) 2019-2021. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package rest

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/micro/go-micro/errors"

	"github.com/pydio/cells/common"
	"github.com/pydio/cells/common/log"
	defaults "github.com/pydio/cells/common/micro"
	"github.com/pydio/cells/common/proto/idm"
	"github.com/pydio/cells/common/proto/tree"
	service "github.com/pydio/cells/common/service/proto"
	"github.com/pydio/cells/common/utils/permissions"
)

type deleteJobs struct {
	RecyclesNodes map[string]*tree.Node
	RecycleMoves  map[string][]string
	Deletes       []string
}

func newDeleteJobs() *deleteJobs {
	return &deleteJobs{
		RecycleMoves:  make(map[string][]string),
		RecyclesNodes: make(map[string]*tree.Node),
	}
}

func sourceInRecycle(ctx context.Context, source *tree.Node, ancestors []*tree.Node) bool {

	for _, n := range ancestors {
		if n.GetStringMeta("name") == common.RecycleBinName {
			return true
		}
	}

	return false
}

func findRecycleForSource(ctx context.Context, source *tree.Node, ancestors []*tree.Node) (recycle *tree.Node, err error) {

	var ids []string
	for _, n := range ancestors {
		if n.Uuid == source.Uuid {
			// Make sure no to use source as recycle root
			continue
		}
		ids = append(ids, n.Uuid)
	}
	q, _ := ptypes.MarshalAny(&idm.ACLSingleQuery{
		NodeIDs: ids,
		Actions: []*idm.ACLAction{permissions.AclRecycleRoot},
	})
	recycleAcls := map[string]bool{}
	cl := idm.NewACLServiceClient(common.ServiceGrpcNamespace_+common.ServiceAcl, defaults.NewClient())
	s, e := cl.SearchACL(ctx, &idm.SearchACLRequest{
		Query: &service.Query{SubQueries: []*any.Any{q}},
	})
	if e != nil {
		err = e
		return
	}
	defer s.Close()
	for {
		r, e := s.Recv()
		if e != nil {
			break
		}
		recycleAcls[r.ACL.NodeID] = true
	}
	for _, n := range ancestors {
		if _, ok := recycleAcls[n.Uuid]; ok {
			recycle = n
			break
		}
	}

	if recycle == nil {
		l := len(ancestors)
		// TODO - for now we just check if it starts with a DATASOURCE:personal
		// But we would need a process that could reverse lookup the template.js
		if l > 3 && strings.HasPrefix(ancestors[l-2].Uuid, "DATASOURCE:personal") {
			personalFolder := ancestors[l-3]
			// This is a "personal files" case, where the recycle root may not have been created for various reasons, add it now
			log.Logger(ctx).Info("Recycle not found inside a personal files, create ACL now on ", ancestors[l-3].Zap())
			newAcl := &idm.ACL{
				NodeID: personalFolder.Uuid,
				Action: permissions.AclRecycleRoot,
			}
			if _, e := cl.CreateACL(ctx, &idm.CreateACLRequest{ACL: newAcl}); e == nil {
				return personalFolder, nil
			}
		}
		err = errors.NotFound("RecycleNotFound", "cannot find recycle root on this branch")
	}
	return
}
