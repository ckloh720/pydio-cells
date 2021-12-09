package service

import (
	"context"

	"github.com/pydio/cells/v4/common/config/runtime"
	"github.com/pydio/cells/v4/common/server/generic"
	servicecontext "github.com/pydio/cells/v4/common/service/context"
)

// WithGeneric adds a http micro service handler to the current service
func WithGeneric(f func(context.Context, *generic.Server) error) ServiceOption {
	return func(o *ServiceOptions) {
		// Making sure the runtime is correct
		if o.Fork && !runtime.IsFork() {
			return
		}

		o.Server = servicecontext.GetServer(o.Context, "generic")
		o.serverStart = func() error {
			var srvg *generic.Server

			o.Server.As(&srvg)

			return f(o.Context, srvg)
		}

		// TODO v4 import wrappers for the server
	}
}

// WithGenericStop adds a http micro service handler to the current service
func WithGenericStop(f func(context.Context, *generic.Server) error) ServiceOption {
	return func(o *ServiceOptions) {
		o.serverStop = func() error {
			var srvg *generic.Server

			o.Server.As(&srvg)

			return f(o.Context, srvg)
		}
	}
}
