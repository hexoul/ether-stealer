package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/hexoul/ether-stealer/crypto"
	"github.com/hexoul/ether-stealer/json"
	"github.com/hexoul/ether-stealer/log"

	"github.com/korovkin/limiter"
)

const (
	httpTimeout   = 10
	urlForBalance = "https://api.infura.io/v1/jsonrpc/mainnet/eth_getBalance?params=[\"%s\",\"latest\"]"
)

var (
	nLimit int
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
		log.Infof("GOTIT from %s !!SECRET!! %x", addr.String(), privkey)
	} else {
		fmt.Printf("FAILED from %s\n", addr.String())
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
		}
	}
}

func main() {
	// Starting message
	var ip string
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
				}
			}
		}
	}
	log.Info("Steal start!!! from ", ip)

	limit := limiter.NewConcurrencyLimiter(nLimit)
	for {
		//for i := 0; i < 1; i++ {
		pubkey, privkey := crypto.GenerateKeyPair()
		addr := crypto.ToAddressFromPubkey(pubkey)
		limit.Execute(func() { steal(addr, privkey) })
	}
	//limit.Wait()
}
