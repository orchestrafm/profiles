package main

import (
	"github.com/orchestrafm/profiles/src/database"
	"github.com/orchestrafm/profiles/src/routing"
	"github.com/spidernest-go/logger"
)

func main() {
	err := database.Connect()
	logger.Error().
		Err(err).
		Msg("MySQL Database could not be attached to.")
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("MySQL Database could not be attached to.")
	}
	database.Synchronize()

	routers.ListenAndServe()
}
