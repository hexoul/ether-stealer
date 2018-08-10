package contract

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethClient *ethclient.Client
)

const (
	urlForJSONRPC = "https://mainnet.infura.io"
)

func init() {
	ethClient, _ = ethclient.Dial(urlForJSONRPC)
}

// SuggestGas fowarded from ethclient
func SuggestGas() (*big.Int, error) {
	return ethClient.SuggestGasPrice(context.Background())
}
