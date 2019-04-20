package logger

import (
	"os"
	"runtime"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var l zerolog.Logger

func init() {
	// Find the platform-specific directory
	appdir := *new(string)
	switch runtime.GOOS {
	case "windows":
		appdir = os.Getenv("APPDATA")
	default:
		appdir = "/var/log"
	}

	// Allocate a new logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	l = zerolog.New(&lumberjack.Logger{
		Filename:   appdir + "/ofm/profiles.log",
		MaxSize:    128,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   true,
	}).With().Timestamp().Logger()

}
