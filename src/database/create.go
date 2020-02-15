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

	id, err := r.LastInsertId()
	if err == nil {
		p.ID = uint64(id)
	}
	return err
}

func (r *ReqList) New() error {
	_, err := db.InsertInto("reqlist").
		Values(r).
		Exec()

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Email could not be inserted into the table.")
	}

	return err
}
