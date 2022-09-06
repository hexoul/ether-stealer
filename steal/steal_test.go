package steal

import (
	"testing"

	"github.com/hexoul/ether-stealer/crypto"
)

const (
	targetAddr = "0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8"
)

func TestSteal(t *testing.T) {
	pubkey, privkey := crypto.GenerateKeyPair()
	addr := crypto.ToAddressFromPubkey(pubkey)
	Steal(addr, privkey)
}
