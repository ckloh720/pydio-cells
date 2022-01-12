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

// Package gateway spins an S3 gateway for serving files using the Amazon S3 protocol.
package gateway

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/pydio/cells/v4/common/server/middleware"

	pydio "github.com/pydio/cells/v4/gateway/data/gw"

	minio "github.com/minio/minio/cmd"
	"go.uber.org/zap"

	"github.com/pydio/cells/v4/common"
	"github.com/pydio/cells/v4/common/log"
	"github.com/pydio/cells/v4/common/plugins"
	serverhttp "github.com/pydio/cells/v4/common/server/http"
	"github.com/pydio/cells/v4/common/service"
	"github.com/pydio/cells/v4/common/utils/net"
	_ "github.com/pydio/cells/v4/gateway/data/gw"
	"github.com/pydio/cells/v4/gateway/data/hooks"
)

type logger struct {
	ctx context.Context
}

func (l *logger) Info(entry interface{}) {
	log.Logger(l.ctx).Info("Minio Gateway", zap.Any("data", entry))
}

func (l *logger) Error(entry interface{}) {
	log.Logger(l.ctx).Error("Minio Gateway", zap.Any("data", entry))
}

func (l *logger) Audit(entry interface{}) {
	log.Auditer(l.ctx).Info("Minio Gateway", zap.Any("data", entry))
}

func init() {

	plugins.Register("main", func(ctx context.Context) {

		port := net.GetAvailablePort()

		service.NewService(
			service.Name(common.ServiceGatewayData),
			service.Context(ctx),
			service.Tag(common.ServiceTagGateway),
			// service.RouterDependencies(),
			service.Description("S3 Gateway to tree service"),
			//service.Port(fmt.Sprintf("%d", port)),
			service.WithHTTP(func(c context.Context, mux *http.ServeMux) error {

				u, _ := url.Parse(fmt.Sprintf("http://localhost:%d", port))
				proxy := httputil.NewSingleHostReverseProxy(u)
				mux.HandleFunc("/io/", func(writer http.ResponseWriter, request *http.Request) {
					proxy.ServeHTTP(writer, request)
				})
				mux.HandleFunc("/data/", func(writer http.ResponseWriter, request *http.Request) {
					proxy.ServeHTTP(writer, request)
				})

				var certFile, keyFile string
				/*
					if config.Get("cert", "http", "ssl").Bool() {
						certFile = config.Get("cert", "http", "certFile").String()
						keyFile = config.Get("cert", "http", "keyFile").String()
					}
				*/

				srv := &gatewayDataServer{
					port:     port,
					certFile: certFile,
					keyFile:  keyFile,
				}
				go srv.Start(c)

				return nil
			}),
		)
	})
}

type gatewayDataServer struct {
	ctx      context.Context
	cancel   context.CancelFunc
	port     int
	certFile string
	keyFile  string
}

func (g *gatewayDataServer) Start(ctx context.Context) error {
	os.Setenv("MINIO_BROWSER", "off")
	os.Setenv("MINIO_ROOT_USER", common.S3GatewayRootUser)
	os.Setenv("MINIO_ROOT_PASSWORD", common.S3GatewayRootPassword)

	minio.HookRegisterGlobalHandler(serverhttp.ContextMiddlewareHandler(middleware.ClientConnIncomingContext(ctx)))
	minio.HookRegisterGlobalHandler(hooks.GetPydioAuthHandlerFunc("gateway"))
	pydio.PydioGateway = &pydio.Pydio{
		RuntimeCtx: ctx,
	}

	params := []string{"minio", "gateway", "pydio", "--address", fmt.Sprintf(":%d", g.port), "--quiet"}
	minio.Main(params)
	/*
		console := &logger{ctx: g.ctx}
		ctx, cancel := context.WithCancel(g.ctx)
		g.cancel = cancel
		gw := &pydio.Pydio{}
		go minio.StartPydioGateway(ctx, gw, fmt.Sprintf(":%d", g.port), "gateway", "gatewaysecret", console, g.certFile, g.keyFile)
	*/
	return nil
}

func (g *gatewayDataServer) Stop() error {
	if g.cancel != nil {
		g.cancel()
	}

	return nil
}
