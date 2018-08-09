package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// GenerateKeyPair that are public and private key
func GenerateKeyPair() (pubkey, privkey []byte) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	pubkey = elliptic.Marshal(secp256k1.S256(), key.X, key.Y)

	privkey = make([]byte, 32)
	blob := key.D.Bytes()
	copy(privkey[32-len(blob):], blob)

	return pubkey, privkey
}

// Sign given message
func Sign(msg, privkey []byte) ([]byte, error) {
	return secp256k1.Sign(msg, privkey)
}

// RecoverPubkey recover public key from signed message
func RecoverPubkey(msg, sig []byte) ([]byte, error) {
	sig[len(sig)-1] %= 4
	return secp256k1.RecoverPubkey(msg, sig)
}

// ToAddressFromPubkey returns ether address from public key
func ToAddressFromPubkey(pubkey []byte) common.Address {
	return common.BytesToAddress(crypto.Keccak256(pubkey[1:])[12:])
}
