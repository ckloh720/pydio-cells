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

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	action       string
	value        string
	roleID       string
	workspaceID  string
	nodeID       string
	actions      []string
	roleIDs      []string
	workspaceIDs []string
	nodeIDs      []string
)

// AclCmd represents the acl command
var AclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Manage access control lists",
	Long: `
DESCRIPTION

  ACLs are managed in a dedicated microservice.

  It is simpler to manage them in the frontend, but you can use this command to create/delete/search ACLs directly.
  ACLs are used to grant permissions to a given node (retrieved by UUID) for a given Role.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AdminCmd.AddCommand(AclCmd)
}
