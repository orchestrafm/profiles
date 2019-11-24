package database

import (
	"github.com/spidernest-go/logger"
)

func (p *Profile) New() error {
	// TODO: Make sure something doesn't already exist in the spot [id, track_id]
	r, err := db.InsertInto("profiles").
		Values(p).
		Exec()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Profile could not be inserted into the table.")
	}

	p.ID, err = r.LastInsertId()
	return err
}
