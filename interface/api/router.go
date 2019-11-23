package api

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type Handler func(ctx Context) (json.Marshaler, error)

type Router interface {
	Group(path string) Router
	GET(path string, handler Handler)
	POST(path string, handler Handler)
	PUT(path string, handler Handler)
	DELETE(path string, handler Handler)
	PATCH(path string, handler Handler)
	HEAD(path string, handler Handler)
	OPTIONS(path string, handler Handler)
	Handler() fasthttp.RequestHandler
}
