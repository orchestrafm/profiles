package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var l zerolog.Logger

func init() {
	// Allocate a new logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	l = zerolog.New(os.Stdout).With().Timestamp().Logger()
}
