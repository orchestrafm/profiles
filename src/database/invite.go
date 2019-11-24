package database

import (
	"database/sql"
	"errors"

	"github.com/spidernest-go/logger"
)

type Invite struct {
	ID     uint64 `db:"id"`
	Code   string `db:"code"`
	Burned bool   `db:"burned"`
}

func BurnInvite(code string) error {
	invites := db.Collection("invites")
	//rs := invites.Find(code) // BUG: This doesn't work, not sure why
	rs := invites.Find("code", code)
	i := *new(Invite)
	err := rs.One(&i)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")

		return err
	}

	if i.Burned == true {
		logger.Warn().
			Msg("Invite Code was already burned.")

		return errors.New("Invite code was already burned.")
	}
	i.Burned = true
	err = rs.Update(i)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invite Code could not be updated from the table.")
	}
	return err
}

func UnburnInvite(code string) error {
	invites := db.Collection("invites")
	rs := invites.Find("code", code)
	i := *new(Invite)
	err := rs.One(&i)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")

		return err
	}

	if i.Burned == false {
		logger.Warn().
			Msg("Invite code has yet to be used.")
		return nil
	}

	i.Burned = false
	err = rs.Update(i)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invite Code could not be updated from the table.")
	}
	return err
}
