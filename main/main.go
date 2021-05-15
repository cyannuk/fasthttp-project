package main

import (
	"os"

	"fasthttp-project/composition"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// init logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05", NoColor: true}) // .Level(zerolog.ErrorLevel)
}

func main() {
	log.Info().Msg("Starting..")

	err := composition.Application()
	if err != nil {
		log.Error().Err(err).Msg("Application")
	}
}