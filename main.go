package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"fasthttp-project/api"
	"fasthttp-project/api/rest"
	"fasthttp-project/domain/repository"
	"fasthttp-project/domain/service"
)

var bindHost, dbConnectionString string

func init() {
	// init logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05", NoColor: true})
	// init flags
	flag.StringVar(&bindHost, "bind", "0.0.0.0:8080", "set bind host")
	flag.StringVar(&dbConnectionString, "connection_string",
		"host=192.168.0.6 user=postgres password=postgres dbname=demo sslmode=disable",
		"db connection string")
	flag.Parse()
}

func main() {
	log.Info().Msg("Starting..")

	// init database
	dataSource, err := repository.NewDataSource(dbConnectionString)
	if err != nil {
		log.Fatal().Err(err)
		return
	}
	defer dataSource.CloseDataSource()

	userService := service.NewUserService(repository.NewUserRepository(dataSource))
	orderService := service.NewOrderService(repository.NewOrderRepository(dataSource))
	s := api.NewServer()
	s.GET("/user_orders", rest.GetUserOrdersHandler(userService))
	ug := s.Group("/users")
	ug.GET("/", rest.GetUsersHandler(userService))
	ug.POST("/", rest.CreateUserHandler(userService))
	og := ug.Group("/:user_id")
	og.GET("/", rest.GetUserHandler(userService))
	og.GET("/orders", rest.GetOrdersHandler(orderService))
	og.GET("/orders/:order_id", rest.GetOrderHandler(orderService))
	if err := s.Start(":8080"); err != nil {
		log.Fatal().Err(err)
	}
}
