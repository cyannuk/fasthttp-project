package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"fasthttp-project/api"
	"fasthttp-project/domain/repository"
	"fasthttp-project/domain/service"
)

var bindHost, dbConnectionString string

func init() {
	// init logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05", NoColor: true}) // .Level(zerolog.ErrorLevel)
	// init flags
	flag.StringVar(&bindHost, "bind", "0.0.0.0:8080", "set bind host")
	flag.StringVar(&dbConnectionString, "connection_string",
		"host=127.0.0.1 user=postgres password=postgres dbname=demo sslmode=disable connect_timeout=2",
		"db connection string")
	flag.Parse()
}

func main() {
	log.Info().Msg("Starting..")

	// init database
	dataSource, err := repository.NewDataSource(dbConnectionString)
	if err != nil {
		log.Error().Err(err).Msg("DataSource")
		return
	}
	defer dataSource.Close()

	application := api.NewApplication(service.NewUserService(repository.NewUserRepository(dataSource)), service.NewOrderService(repository.NewOrderRepository(dataSource)))
	application.Start(bindHost)
}
