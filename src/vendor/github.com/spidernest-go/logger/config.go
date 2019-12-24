package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

// Configure takes logging parameters and allows custom parameters.
// file: 		Filepath to where the logging file is saved.
// maxsize: 	Size of the file allowed in megabytes.
// maxbackups: 	Maximum allowed backups of a single log file.
// maxage:		Maximum allowed age of a log file.
// compress:	Compress the logfile?
// timefield: 	zerolog time format
func Configure(file string, maxsize, maxbackups, maxage int, compress bool, timefield string) {
	zerolog.TimeFieldFormat = timefield
	l = zerolog.New(&lumberjack.Logger{
		Filename:   file,
		MaxSize:    maxsize,
		MaxBackups: maxbackups,
		MaxAge:     maxage,
		Compress:   compress,
	}).With().Timestamp().Logger()
}

func Logger() zerolog.Logger {
	return l
}

func Stdout() {
	l = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()
}
