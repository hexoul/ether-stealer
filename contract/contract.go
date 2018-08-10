package contract

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/hexoul/ether-stealer/contract/abigen/omg"
)

// ERC20 token structure
type ERC20 struct {
	addr      string
	name      string
	unit      string
	contract  interface{}
	balanceOf func(common.Address) (*big.Int, error)
}

var (
	ethClient *ethclient.Client
	tokens    []ERC20
)

const (
	urlForJSONRPC = "https://mainnet.infura.io"
)

func init() {
	ethClient, _ = ethclient.Dial(urlForJSONRPC)

	// Initialize tokens info
	tokens = append(tokens, ERC20{addr: "0xd26114cd6EE289AccF82350c8d8487fedB8A0C07", name: "OmiseGO", unit: "OMG"})

	// Set caller
	for i, token := range tokens {
		switch token.unit {
		case "OMG":
			caller, _ := omg.NewOMGTokenCaller(common.HexToAddress(token.addr), ethClient)
			token.contract = caller
			token.balanceOf = func(addr common.Address) (*big.Int, error) {
				return caller.BalanceOf(&bind.CallOpts{}, addr)
			}
		default:
			continue
		}
		tokens[i] = token
	}
}

// SuggestGas fowarded from ethclient
func SuggestGas() (*big.Int, error) {
	return ethClient.SuggestGasPrice(context.Background())
}

// List ERC20 tokens registered
func List() []ERC20 {
	return tokens
}

// CanSteal ERC20 token of this address
func CanSteal(addr common.Address) (ret string) {
	for _, token := range tokens {
		if val, err := token.balanceOf(addr); err == nil {
			if val.Uint64() > 0 {
				ret += token.unit + " "
			}
		}
	}
	return
}
