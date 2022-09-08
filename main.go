package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/hexoul/ether-stealer/crypto"
	"github.com/hexoul/ether-stealer/steal"

	"github.com/korovkin/limiter"
)

var (
	nLimit     *int
	identifier *string
)

func init() {
	nLimit = flag.Int("concurrency", 10, "The number of threads can be executed concurrently.")
	identifier = flag.String("id", "", "An identifier of a client.")
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

	flag.Parse()
	stealer := steal.New()
	fmt.Printf("Start to steal!!! from <%s> by <%s>\n", ip, *identifier)

	limit := limiter.NewConcurrencyLimiter(*nLimit)
	// If you want finite iterator,
	// 1. Add a condition to `for` statement. e.g. `for i:=0; i<N`
	// 2. Put `limit.WaitAndClose()` after `for` statement.
	for {
		pubkey, privkey := crypto.GenerateKeyPair()
		addr := crypto.PubkeyToAddress(pubkey)
		limit.Execute(func() { stealer.Steal(addr, privkey) })
	}
	// limit.WaitAndClose()
}
