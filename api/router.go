package api

import (
	httprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"gopkg.in/reform.v1"

	"fasthttp-project/api/errors"
	"fasthttp-project/interface/api"
)

type router struct {
	*httprouter.Router
}

func defaultHandler(handler api.Handler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		entity, err := handler(context{ctx})
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

func (r router) GET(path string, handler api.Handler) {
	r.Router.GET(path, defaultHandler(handler))
}

func (r router) POST(path string, handler api.Handler) {
	r.Router.POST(path, defaultHandler(handler))
}

func (r router) PUT(path string, handler api.Handler) {
	r.Router.PUT(path, defaultHandler(handler))
}

func (r router) DELETE(path string, handler api.Handler) {
	r.Router.DELETE(path, defaultHandler(handler))
}

func (r router) PATCH(path string, handler api.Handler) {
	r.Router.PATCH(path, defaultHandler(handler))
}

func (r router) HEAD(path string, handler api.Handler) {
	r.Router.HEAD(path, defaultHandler(handler))
}

func (r router) OPTIONS(path string, handler api.Handler) {
	r.Router.OPTIONS(path, defaultHandler(handler))
}

func (r router) Group(path string) router {
	return router{r.Router.Group(path)}
}

func (r router) Handler() fasthttp.RequestHandler {
	return r.Router.Handler
}

func NewRouter() router {
	return router{httprouter.New()}
}
