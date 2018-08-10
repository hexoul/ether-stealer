package contract

import (
	"testing"
)

func TestSuggestGas(t *testing.T) {
	if _, err := SuggestGas(); err != nil {
		t.Fatal("Failed SuggestGas")
	}
}
