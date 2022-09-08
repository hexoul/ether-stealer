package crypto

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	testmsg  = hexutil.MustDecode("0xce0677bb30baa8cf067c88db9811f4333d131bf8bcf12fe7065d211dce971008")
	testpriv = hexutil.MustDecode("0x25c317c8d0a63c122073ae52984e8477e7fbc322c93a9457c5579ee6e5a813b3")
	testaddr = hexutil.MustDecode("0xeD56062123B0301A9a642f85F2711581bEc8D79D")
)

func TestGenerateKey(t *testing.T) {
	pub, priv := GenerateKeyPair()
	t.Logf("\n%x\n%x", pub, priv)
}

func TestSignAndRecover(t *testing.T) {
	pub, priv := GenerateKeyPair()
	sig, err := Sign(testmsg, priv)
	if err != nil {
		t.Fatal("Failed to sign.")
	}
	rpub, err := RecoverPubkey(testmsg, sig)
	if err != nil {
		t.Fatal("Failed to recover.")
	} else if !bytes.Equal(pub, rpub) {
		t.Fatalf("Mismatch btw %x and %x", pub, rpub)
	}
}

func TestPubkeyToAddress(t *testing.T) {
	sig, err := Sign(testmsg, testpriv)
	if err != nil {
		t.Fatal("Failed to sign.")
	}
	rpub, err := RecoverPubkey(testmsg, sig)
	if err != nil {
		t.Fatal("Failed to recover.")
	}
	addr := PubkeyToAddress(rpub)
	t.Logf("%x", addr)
	if !bytes.Equal(addr.Bytes(), testaddr) {
		t.Fatal("Failed to get an address from given pubkey.")
	}
}

func TestPrivkeyToAddress(t *testing.T) {
	addr, err := PrivkeyToAddress(testpriv)
	if err != nil {
		t.Fatal("Failed to get an address from given privkey.")
	}
	t.Logf("%x", addr)
}
