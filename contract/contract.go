package contract

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/hexoul/ether-stealer/contract/abigen/bat"
	"github.com/hexoul/ether-stealer/contract/abigen/npxs"
	"github.com/hexoul/ether-stealer/contract/abigen/omg"
	"github.com/hexoul/ether-stealer/contract/abigen/zrx"
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
	tokens = append(tokens, ERC20{addr: "0xA15C7Ebe1f07CaF6bFF097D8a589fb8AC49Ae5B3", name: "Pundi X Token", unit: "NPXS"})
	tokens = append(tokens, ERC20{addr: "0xE41d2489571d322189246DaFA5ebDe1F4699F498", name: "ZeroEx", unit: "ZRX"})
	tokens = append(tokens, ERC20{addr: "0x0D8775F648430679A709E98d2b0Cb6250d2887EF", name: "Basic Attention Token", unit: "BAT"})

	// Set caller
	for i, token := range tokens {
		switch token.unit {
		case "OMG":
			if caller, err := omg.NewOMGTokenCaller(common.HexToAddress(token.addr), ethClient); err == nil {
				token.contract = caller
				token.balanceOf = func(addr common.Address) (*big.Int, error) {
					return caller.BalanceOf(&bind.CallOpts{}, addr)
				}
			}
			break
		case "NPXS":
			if caller, err := npxs.NewNPXSTokenCaller(common.HexToAddress(token.addr), ethClient); err == nil {
				token.contract = caller
				token.balanceOf = func(addr common.Address) (*big.Int, error) {
					return caller.BalanceOf(&bind.CallOpts{}, addr)
				}
			}
			break
		case "ZRX":
			if caller, err := zrx.NewZRXTokenCaller(common.HexToAddress(token.addr), ethClient); err == nil {
				token.contract = caller
				token.balanceOf = func(addr common.Address) (*big.Int, error) {
					return caller.BalanceOf(&bind.CallOpts{}, addr)
				}
			}
			break
		case "BAT":
			if caller, err := bat.NewBATokenCaller(common.HexToAddress(token.addr), ethClient); err == nil {
				token.contract = caller
				token.balanceOf = func(addr common.Address) (*big.Int, error) {
					return caller.BalanceOf(&bind.CallOpts{}, addr)
				}
			}
			break
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
		if token.contract == nil {
			continue
		}
		if val, err := token.balanceOf(addr); err == nil {
			if val.Uint64() > 0 {
				ret += token.unit + " "
			}
		}
	}
	return
}
