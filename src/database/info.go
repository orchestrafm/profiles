package database

import (
	"strconv"

	"github.com/spidernest-go/logger"
)

func SelectProfileById(id uint64) (error, *Profile) {
	pf := *new(Profile)
	err := db.SelectFrom("profiles").
		Where("id = " + strconv.FormatUint(id, 10)).
		Limit(1).
		One(&pf)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("SQL Execution had an issue when executing.")
		return err, nil
	}

	// pf.UUID = *new(string) // QUEST: Should I clear this field so it doesn't return?

	return nil, &pf
}
