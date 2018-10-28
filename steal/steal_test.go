package steal

import (
	"testing"
)

const (
	targetAddr = "0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8"
)

func TestHasBalance(t *testing.T) {
	if b, _ := hasBalance(targetAddr); !b {
		t.Error("Failed to check balance")
	}
}

func TestHasTxCount(t *testing.T) {
	if b, _ := hasTxCount(targetAddr); !b {
		t.Error("Failed to check TX count")
	}
}
