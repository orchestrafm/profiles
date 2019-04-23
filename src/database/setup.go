package database

import (
	"database/sql"
	"os"
	"strings"

	"github.com/orchestrafm/profiles/src/logger"
	"github.com/orchestrafm/profiles/src/static"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

const version uint8 = 1

var db sqlbuilder.Database

func Connect() error {
	opts := make(map[string]string)
	opts["parseTime"] = "True"
	conn := mysql.ConnectionURL{
		Database: os.Getenv("MYSQL_DB"),
		Host:     os.Getenv("MYSQL_HOST"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASS"),
		Options:  opts,
	}

	err := *new(error)
	db, err = mysql.Open(conn)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("MySQL Database Unreachable.")

		return err
	}

	return nil
}

func UpgradeIfOutdated() {
	err, exists := doesTableExist("meta")
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("SQL Query did not execute.")
	}

	if exists == false {
		logger.Warn().
			Err(err).
			Msg("Meta Information Table does not exist, creating it...")

	}

	switch version {
	case 0:
		update("001_table.sql")
		fallthrough

	default:
		db.Exec("UPDATE meta SET version = ? WHERE tablename = ?", version, "profiles")
		logger.Info().
			Msg("All schema upgrades finished.")
	}
}

func doesTableExist(name string) (error, bool) {
	stmt, err := db.Prepare(`SELECT *
        FROM information_schema.tables
        WHERE table_schema = ?
            AND table_name = ?
        LIMIT 1;`)

	if err != nil {
		return err, true
	}

	rows := stmt.QueryRow(os.Getenv("MYSQL_DB"), name)

	table := make(map[string]interface{})
	err = rows.Scan(&table)

	if err == sql.ErrNoRows {
		return nil, false
	}

	//TODO: this should be it's own error through errors.New but I'm lazy
	return nil, true
}

func update(filename string) {
	f, err := static.ReadFile(filename)

	if err != nil {
		logger.Fatal().
			Err(err).
			Msgf("Static Resource file could not be opened or found (%s).", filename)
	}

	buf := new(strings.Builder)
	buf.Write(f)

	stmt, err := db.Prepare(buf.String())

	if err != nil {
		logger.Fatal().
			Err(err).
			Msgf("SQL Prepare Statement failed (%s).", buf.String())
	}

	_, err = stmt.Exec()

	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("SQL Statement failed to execute.")
	}
}
