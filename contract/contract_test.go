package contract

import (
	"testing"
)

func TestSuggestGas(t *testing.T) {
	if _, err := SuggestGas(); err != nil {
		t.Fatal("Failed SuggestGas")
	}
}

func TestList(t *testing.T) {
	tokens := List()
	for _, token := range tokens {
		t.Logf("%v", token)
	}
}
