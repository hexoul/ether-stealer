package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHasBalance(t *testing.T) {
	if b, _ := hasBalance("0xc94770007dda54cF92009BFF0dE90c06F603a09f"); !b {
		t.Error("Failed to get balance")
	}
}

func TestHasTxCount(t *testing.T) {
	if b, _ := hasTxCount("0x3f5ce5fbfe3e9af3971dd833d26ba9b5c936f0be"); !b {
		t.Error("Failed to get TxCount")
	}
}
