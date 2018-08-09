package log

import (
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	Info("ether", "stealer")
	Infof("telegram test %s", "lol")
	time.Sleep(2 * time.Second)
}
