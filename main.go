package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/hexoul/ether-stealer/json"

	_ "github.com/ethereum/go-ethereum/crypto"
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

func isBalanceGreaterThanZero(addr string) bool {
	url := fmt.Sprintf(urlForBalance, addr)
	if resp, err := httpClient.Get(url); err == nil {
		if respBody, err := ioutil.ReadAll(resp.Body); err == nil {
			rpcResp := json.GetRPCResponseFromJSON(respBody)
			fmt.Print(rpcResp.String())
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Hello")
}
