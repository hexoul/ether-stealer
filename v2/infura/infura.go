package infura

import (
	"flag"
	"fmt"

	"github.com/hexoul/ether-stealer/v2/json"
)

var (
	apiKey *string
)

func init() {
	apiKey = flag.String("infura-apikey", "", "API key of your Infura project.")
}

type Infura struct {
	apiKey  string
	baseUrl string
}

func New() *Infura {
	if *apiKey == "" {
		panic("Infura API key is required.")
	}

	return &Infura{
		apiKey:  *apiKey,
		baseUrl: fmt.Sprintf("https://mainnet.infura.io/v3/%s", *apiKey),
	}
}

func (i *Infura) HasBalance(addr string) (b bool, val string) {
	return i.get("eth_getBalance", []string{addr, "latest"})
}

func (i *Infura) get(method string, params []string) (b bool, val string) {
	rpcRequest := json.RPCRequest{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}

	ret, err := json.GetRPCResponse(i.baseUrl, rpcRequest)
	if err != nil {
		fmt.Println(fmt.Sprintf("infura: error occurred. <%s>", err.Error()))
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
