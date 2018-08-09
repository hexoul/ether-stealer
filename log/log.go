package log

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	logger *log.Logger
)

func init() {
	// Initalize logger
	logger = log.New()

	// Default configuration
	timestampFormat := "02-01-2006 15:04:05"
	logger.Formatter = &log.TextFormatter{
		TimestampFormat: timestampFormat,
		FullTimestamp:   true,
	}
	if f, err := os.OpenFile("./stealer.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err == nil {
		logger.Out = io.MultiWriter(f, os.Stdout)
	} else {
		fmt.Print("Failed to open log file: you can miss important log")
		logger.Out = os.Stdout
	}
	logger.SetLevel(log.InfoLevel)
}

// Info level logging
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof info-level logging with format
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}
