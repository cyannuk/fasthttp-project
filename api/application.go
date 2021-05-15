package api

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"fasthttp-project/interface/service"
	"github.com/fasthttp/router"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type Application struct {
	server *fasthttp.Server
	router *router.Router
	userService  service.UserService
	orderService service.OrderService
}

func (application *Application) Start(addr string) {
	if err := application.server.ListenAndServe(addr); err != nil {
		log.Error().Err(err).Msg("Application")
	}
}

func (application *Application) Stop() {
	log.Info().Msg("Exiting..")
	err := application.server.Shutdown()
	if err != nil {
		log.Error().Err(err).Msg("Application")
	}
}

func NewApplication(userService service.UserService, orderService service.OrderService) Application {
	application := Application{&fasthttp.Server{NoDefaultServerHeader: true}, router.New(), userService, orderService}

	once := sync.Once{}
	channel := make(chan os.Signal, 3)
	signal.Notify(channel, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		<-channel
		once.Do(application.Stop)
	}()

	application.router.Handle(fasthttp.MethodGet, "/users", application.getUsers)
	application.router.Handle(fasthttp.MethodPost, "/users", application.createUser)

	application.router.Handle(fasthttp.MethodGet, "/users/{user_id}", application.getUser)
	application.router.Handle(fasthttp.MethodGet, "/users/{user_id}/orders", application.getOrders)
	application.router.Handle(fasthttp.MethodGet, "/users/{user_id}/orders/{order_id}", application.getUserOrder)

	application.router.Handle(fasthttp.MethodGet, "/user_orders", application.getUserOrders)

	application.server.Handler = application.router.Handler;
	return application
}
