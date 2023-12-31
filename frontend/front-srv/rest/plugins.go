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

// Package rest is a service for serving specific requests directly to frontend
package rest

import (
	"context"
	"encoding/gob"
	"os"

	"github.com/pydio/cells/common/config"

	"github.com/pydio/packr"

	"github.com/pydio/cells/common"
	"github.com/pydio/cells/common/plugins"
	"github.com/pydio/cells/common/service"
	"github.com/pydio/cells/common/service/frontend"
	"github.com/pydio/cells/frontend/front-srv/rest/modifiers"
)

var BasePluginsBox = frontend.PluginBox{
	Box: packr.NewBox("../../../frontend/front-srv/assets"),
	Exposes: []string{
		"access.gateway",
		"access.homepage",
		"access.settings",
		"action.avatar",
		"action.compression",
		"action.migration",
		"action.share",
		"action.user",
		"auth.pydio",
		"authfront.session_login",
		"core.activitystreams",
		"core.auth",
		"core.authfront",
		"core.conf",
		"core.mailer",
		"core.pydio",
		"core.uploader",
		"editor.browser",
		"editor.ckeditor",
		"editor.codemirror",
		"editor.diaporama",
		"editor.exif",
		"editor.infopanel",
		"editor.libreoffice",
		"editor.openlayer",
		"editor.pdfjs",
		"editor.soundmanager",
		"editor.text",
		"editor.video",
		"gui.ajax",
		"gui.mobile",
		"meta.comments",
		"meta.exif",
		"meta.simple_lock",
		"meta.user",
		"meta.versions",
		"uploader.html",
		"uploader.http",
	},
}

func init() {

	if os.Getenv("CELLS_ENABLE_LIVEKIT") != "" {
		BasePluginsBox.Exposes = append(BasePluginsBox.Exposes, "action.livekit")
	}
	config.RegisterVaultKey("frontend/plugin/action.livekit/LK_API_SECRET")

	plugins.Register("main", func(ctx context.Context) {
		gob.Register(map[string]string{})

		frontend.RegisterRegModifier(modifiers.MetaUserRegModifier)
		frontend.RegisterPluginModifier(modifiers.MobileRegModifier)

		frontend.WrapAuthMiddleware(modifiers.LogoutAuth)
		frontend.WrapAuthMiddleware(modifiers.RefreshAuth)

		frontend.WrapAuthMiddleware(modifiers.LoginPasswordAuth)
		frontend.WrapAuthMiddleware(modifiers.LoginExternalAuth)
		frontend.WrapAuthMiddleware(modifiers.AuthorizationCodeAuth)

		frontend.WrapAuthMiddleware(modifiers.LoginSuccessWrapper)
		frontend.WrapAuthMiddleware(modifiers.LoginFailedWrapper)

		s := service.NewService(
			service.Name(common.ServiceRestNamespace_+common.ServiceFrontend),
			service.Context(ctx),
			service.Tag(common.ServiceTagFrontend),
			service.Description("REST service for serving specific requests directly to frontend"),
			service.PluginBoxes(BasePluginsBox),
			service.WithWeb(func() service.WebHandler {
				return NewFrontendHandler()
			}),
		)
		// Make sure to have the WebSession wrapper happen before the policies
		// Exclude POST binaries for using Cookies as it's the only one subject to possible CSRF
		s.Init(service.WithWebSession("POST:/frontend/binaries"))
	})

	if os.Getenv("CELLS_ENABLE_FORMS_DEVEL") == "1" {
		config.RegisterExposedConfigs("pydio.rest.forms-devel", formDevelConfigs)
	}

}
