package main

import (
	"time"

	"github.com/orchestrafm/profiles/src/database"
	"github.com/orchestrafm/profiles/src/identity"
	"github.com/orchestrafm/profiles/src/routing"
	"github.com/spidernest-go/logger"
)

func main() {
	err := database.Connect()
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("MySQL Database could not be attached to.")
	}
	database.Synchronize()

	identity.Handshake()
	job := time.NewTicker(time.Minute)
	go func() {
		for {
			<-job.C
			identity.Handshake()
		}
	}()

	routers.ListenAndServe()
}
