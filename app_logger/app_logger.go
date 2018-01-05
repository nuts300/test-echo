package appLogger

import (
	"log"
	"os"
)

type (
	CustomLogger interface {
		Error(...interface{})
		Info(...interface{})
		Warn(...interface{})
		Debug(...interface{})
	}

	customeLogger struct {
	}
)

var errorLogger = log.New(os.Stdout, "[error]", log.LstdFlags|log.Llongfile)
var infoLogger = log.New(os.Stdout, "[info]", 0)
var warnLogger = log.New(os.Stdout, "[warn]", 0)
var debugLogger = log.New(os.Stdout, "[debug]", log.LstdFlags|log.Llongfile)

func (c *customeLogger) Error(args ...interface{}) {
	errorLogger.Println(args)
}

func (c *customeLogger) Info(args ...interface{}) {
	infoLogger.Println(args)
}

func (c *customeLogger) Warn(args ...interface{}) {
	warnLogger.Println(args)
}

func (c *customeLogger) Debug(args ...interface{}) {
	debugLogger.Println(args)
}

var logger = customeLogger{}

func GetLogger() CustomLogger {
	return &logger
}
