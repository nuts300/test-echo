package appLogger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type (
	CustomLogger interface {
		Error(...interface{})
		Info(...interface{})
		Warn(...interface{})
		Debug(...interface{})
		Fatal(...interface{})
	}

	customeLogger struct {
		logger *log.Logger
	}
)

func loggerNew() CustomLogger {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	return &customeLogger{
		logger: log.New(),
	}
}

var logger = loggerNew()

func (l *customeLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *customeLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *customeLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *customeLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *customeLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func GetLogger() CustomLogger {
	return logger
}
