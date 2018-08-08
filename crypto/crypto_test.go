package crypto

import (
	"testing"
)

func TestGenerateKey(t *testing.T) {
	pub, priv := GenerateKeyPair()
	t.Logf("\n%x\n%x", pub, priv)
}
