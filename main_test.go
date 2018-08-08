package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestIsBalance(t *testing.T) {
	if !isBalanceGreaterThanZero("0xc94770007dda54cF92009BFF0dE90c06F603a09f") {
		t.Error("Failed to get balance")
	}
}
