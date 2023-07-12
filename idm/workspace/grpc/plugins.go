/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
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

// Package grpc provides the persistence for workspaces
package grpc

import (
	"context"
	"github.com/pydio/cells/v4/common/dao/mysql"
	"github.com/pydio/cells/v4/common/dao/sqlite"
	commonsql "github.com/pydio/cells/v4/common/sql"

	"google.golang.org/grpc"

	"github.com/pydio/cells/v4/common"
	"github.com/pydio/cells/v4/common/broker"
	"github.com/pydio/cells/v4/common/proto/idm"
	"github.com/pydio/cells/v4/common/runtime"
	"github.com/pydio/cells/v4/common/service"
	"github.com/pydio/cells/v4/common/service/resources"
	sqlresources "github.com/pydio/cells/v4/common/sql/resources"
	"github.com/pydio/cells/v4/idm/workspace"
)

const (
	ServiceName = common.ServiceGrpcNamespace_ + common.ServiceWorkspace
)

func init() {
	runtime.Register("main", func(ctx context.Context) {
		var s service.Service
		s = service.NewService(
			service.Name(ServiceName),
			service.Context(ctx),
			service.Tag(common.ServiceTagIdm),
			service.Description("Workspaces Service"),
			service.WithTODOStorage(workspace.NewDAO, commonsql.NewDAO,
				service.WithStoragePrefix("idm_workspace"),
				service.WithStorageSupport(mysql.Driver, sqlite.Driver),
			),
			// service.WithStorage(workspace.NewDAO, service.WithStoragePrefix("idm_workspace")),
			service.WithGRPC(func(ctx context.Context, server grpc.ServiceRegistrar) error {

				h := NewHandler(ctx, s)
				idm.RegisterWorkspaceServiceServer(server, h)

				// Register a cleaner for removing a workspace when there are no more ACLs on it.
				wsCleaner := NewWsCleaner(ctx, h)
				cleaner := &resources.PoliciesCleaner{
					DAO: service.DAOProvider[sqlresources.DAO](s),
					Options: resources.PoliciesCleanerOptions{
						SubscribeRoles: true,
						SubscribeUsers: true,
					},
					LogCtx: ctx,
				}
				if e := broker.SubscribeCancellable(ctx, common.TopicIdmEvent, func(ctx context.Context, message broker.Message) error {
					ev := &idm.ChangeEvent{}
					if e := message.Unmarshal(ev); e == nil {
						_ = wsCleaner.Handle(ctx, ev)
						return cleaner.Handle(ctx, ev)
					}
					return nil
				}); e != nil {
					return e
				}
				return nil
			}),
		)
	})
}
