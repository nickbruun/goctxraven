package ctxraven

import (
	"github.com/getsentry/raven-go"
	"golang.org/x/net/context"

	"github.com/nickbruun/goctxerror"
	"github.com/nickbruun/goctxreq"
)

// Raven error handler.
func ravenErrorHandler(ctx context.Context, err error, msg string) {
	extra := []raven.Interface{
		raven.NewException(err, raven.NewStacktrace(3, 3, nil)),
	}

	if req, ok := ctxreq.FromContext(ctx); ok {
		extra = append(extra, raven.NewHttp(req))
	}

	raven.Capture(raven.NewPacket(msg, extra...), nil)
}

// New context with Raven error handling assigned.
func NewContext(ctx context.Context) context.Context {
	return ctxerror.NewContext(ctx, ravenErrorHandler)
}
