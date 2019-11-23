package api

import (
	"github.com/valyala/fasthttp"

	"fasthttp-project/interface/api"
)

type server struct {
	*fasthttp.Server
	router router
}

func (s server) Start(addr string) error {
	return s.Server.ListenAndServe(addr)
}

func (s server) Group(path string) router {
	return s.router.Group(path)
}

func (s server) GET(path string, handler api.Handler) {
	s.router.GET(path, handler)
}

func (s server) POST(path string, handler api.Handler) {
	s.router.POST(path, handler)
}

func (s server) PUT(path string, handler api.Handler) {
	s.router.PUT(path, handler)
}

func (s server) DELETE(path string, handler api.Handler) {
	s.router.DELETE(path, handler)
}

func (s server) PATCH(path string, handler api.Handler) {
	s.router.PATCH(path, handler)
}

func (s server) HEAD(path string, handler api.Handler) {
	s.router.HEAD(path, handler)
}

func (s server) OPTIONS(path string, handler api.Handler) {
	s.router.OPTIONS(path, handler)
}

func NewServer() server {
	router := NewRouter()
	return server{&fasthttp.Server{Handler: router.Handler(), NoDefaultServerHeader: true}, router}
}
