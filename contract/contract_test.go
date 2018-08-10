package contract

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func TestOmg(t *testing.T) {
	tokens := List()
	for _, token := range tokens {
		if token.unit == "OMG" {
			val, err := token.balanceOf(common.HexToAddress("0x5e44c3e467a49c9ca0296a9f130fc433041aaa28"))
			t.Logf("%d %v", val.Uint64(), err)
		}
	}
}
