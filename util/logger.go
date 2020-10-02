package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type (
	Logger struct {
		Err *log.Logger
		Out *log.Logger
	}
)

func NewLogger() *Logger {
	return &Logger{
		Out: &log.Logger{
			Formatter: new(log.JSONFormatter),
			Out:       os.Stdout,
			Level:     log.InfoLevel,
		},
		Err: &log.Logger{
			Formatter: new(log.JSONFormatter),
			Out:       os.Stderr,
			Level:     log.InfoLevel,
		},
	}
}
