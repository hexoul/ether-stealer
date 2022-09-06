package infura

import (
	"testing"
)

const (
	targetAddr = "0xDA9dfA130Df4dE4673b89022EE50ff26f6EA73Cf"
)

func TestHasBalance(t *testing.T) {
	if b, _ := New().HasBalance(targetAddr); !b {
		t.Error("Failed to check balance")
	}
}
