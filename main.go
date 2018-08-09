package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/hexoul/ether-stealer/crypto"
	"github.com/hexoul/ether-stealer/json"
)

const (
	httpTimeout   = 10
	urlForBalance = "https://api.infura.io/v1/jsonrpc/mainnet/eth_getBalance?params=[\"%s\",\"latest\"]"
)

var (
	httpClient *http.Client
)

func init() {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Second * httpTimeout,
		}).Dial,
		TLSHandshakeTimeout: time.Second * httpTimeout,
	}
	httpClient = &http.Client{
		Timeout:   time.Second * httpTimeout,
		Transport: netTransport,
	}
}

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

func main() {
	pub, priv := crypto.GenerateKeyPair()
	fmt.Printf("\n%x\n%x\n", pub, priv)
	addr := crypto.ToAddressFromPubkey(pub)
	fmt.Printf("%x\n", addr)
}
