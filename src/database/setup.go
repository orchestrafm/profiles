package database

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/orchestrafm/profiles/src/static"
	"github.com/spidernest-go/db/lib/sqlbuilder"
	"github.com/spidernest-go/db/mysql"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/migrate"
)

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

func Synchronize() {
	// TODO: We can get names from reading fileb0x.toml instead, automating this
	names := []string{"000_table.sql"}
	versions := *new([]uint8)
	times := *new([]time.Time)
	buffers := *new([]io.Reader)

	for i := range names {
		// Assign Times
		finfo, err := static.FS.Stat(static.CTX, names[i])
		if err != nil {
			logger.Panic().
				Err(err).
				Msg("Embedded file, " + names[i] + ", does not exists.")
		}
		times = append(times, finfo.ModTime())

		// Assign Versioning
		ver, err := strconv.Atoi(names[i][0:3])
		if err != nil {
			logger.Panic().
				Err(err).
				Msg("Embedded file, " + names[i] + ", could not properly convert it's prefix to a version number.")
		}
		versions = append(versions, uint8(ver))

		// Assign Readers
		data, err := static.ReadFile(names[i])
		if err != nil {
			logger.Panic().
				Err(err).
				Msg("Embedded file, " + names[i] + ", does not exist, or could not be read.")
		}
		buf := bytes.NewBuffer(data)
		buffers = append(buffers, buf)
	}

	if err := migrate.UpTo(versions, names, times, buffers, db); err != nil {
		logger.Panic().
			Err(err).
			Msg("Database Synchronization was unable to complete.")
	}

	logger.Info().
		Msg("Database Synchronization completed successfully.")
}
