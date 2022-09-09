package log

import (
	"flag"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	flag.Parse()
	logger := New()
	logger.Info("Ether", "stealer")
	logger.Infof("Telegram messaging test! %s", "LOL")
	time.Sleep(2 * time.Second)
}
