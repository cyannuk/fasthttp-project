package api

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type HttpServer struct {
	*fasthttp.Server
}

func onShutdown(f func()) {
	once := &sync.Once{}
	signalsChannel := make(chan os.Signal, 3)
	signal.Notify(signalsChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		<-signalsChannel
		once.Do(f)
	}()
}

func NewHttpServer(router RequestRouter) HttpServer {
	s := HttpServer{&fasthttp.Server{Handler: router.Handler, NoDefaultServerHeader: true}}
	onShutdown(func() {
		err := s.Shutdown()
		if err != nil {
			log.Error().Err(err).Msg("HttpServer")
		}
	})
	return s
}
