package api

import (
	"encoding/json"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"gopkg.in/reform.v1"

	"fasthttp-project/api/errors"
)

type Handler func(ctx Context) (json.Marshaler, error)

type RequestRouter struct {
	*router.Router
}

func adaptHandler(handler Handler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		entity, err := handler(Context{ctx})
		if err == nil {
			if entity != nil {
				messageBytes, err := entity.MarshalJSON()
				if err != nil {
					ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
				}
				ctx.SetContentType("application/json")
				ctx.Response.AppendBody(messageBytes)
			}
		} else {
			var statusCode int
			switch err {
			case errors.ErrLimit, errors.ErrOffset, errors.ErrUserId, errors.ErrOrderId:
				statusCode = fasthttp.StatusBadRequest
			case errors.ErrForbidden:
				statusCode = fasthttp.StatusForbidden
			case reform.ErrNoRows:
				statusCode = fasthttp.StatusNotFound
			default:
				statusCode = fasthttp.StatusInternalServerError
			}
			ctx.Error(err.Error(), statusCode)
		}
	}
}

func (r RequestRouter) GET(path string, handler Handler) {
	r.Router.GET(path, adaptHandler(handler))
}

func (r RequestRouter) POST(path string, handler Handler) {
	r.Router.POST(path, adaptHandler(handler))
}

func (r RequestRouter) PUT(path string, handler Handler) {
	r.Router.PUT(path, adaptHandler(handler))
}

func (r RequestRouter) DELETE(path string, handler Handler) {
	r.Router.DELETE(path, adaptHandler(handler))
}

func (r RequestRouter) PATCH(path string, handler Handler) {
	r.Router.PATCH(path, adaptHandler(handler))
}

func (r RequestRouter) HEAD(path string, handler Handler) {
	r.Router.HEAD(path, adaptHandler(handler))
}

func (r RequestRouter) OPTIONS(path string, handler Handler) {
	r.Router.OPTIONS(path, adaptHandler(handler))
}

func (r RequestRouter) Group(path string) RequestRouter {
	return RequestRouter{r.Router.Group(path)}
}

func NewRequestRouter() RequestRouter {
	return RequestRouter{router.New()}
}
