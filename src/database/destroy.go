package database

import (
	"github.com/spidernest-go/logger"
)

func Remove(email string) error {
	err := db.Collection("reqlist").
		Find(email).Delete()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Entry did not exist or could not be deleted.")
	}
	return err
}
