package logger

import (
	"context"

	"github.com/rs/zerolog"
)

func Debug() *zerolog.Event {
	return l.Debug()
}

func Err(err error) *zerolog.Event {
	return l.Err(err)
}

func Error() *zerolog.Event {
	return l.Error()
}

func Fatal() *zerolog.Event {
	return l.Fatal()
}

func Info() *zerolog.Event {
	return l.Info()
}

func Log() *zerolog.Event {
	return l.Log()
}

func Panic() *zerolog.Event {
	return l.Panic()
}

func Print(v ...interface{}) {
	l.Print(v...)
}

func Printf(format string, v ...interface{}) {
	l.Printf(format, v...)
}

func UpdateContext(update func(c zerolog.Context) zerolog.Context) {
	l.UpdateContext(update)
}

func Warn() *zerolog.Event {
	return l.Warn()
}

func With() zerolog.Context {
	return l.With()
}

func WithContext(ctx context.Context) context.Context {
	return l.WithContext(ctx)
}

func WithLevel(level zerolog.Level) *zerolog.Event {
	return l.WithLevel(level)
}

func Write(p []byte) (n int, err error) {
	return l.Write(p)
}
