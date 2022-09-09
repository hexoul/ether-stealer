package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/hexoul/ether-stealer/v2/crypto"
	"github.com/hexoul/ether-stealer/v2/steal"

	"github.com/korovkin/limiter"
)

var (
	concurrency *int
	iteration   *int
	identifier  *string
)

func init() {
	concurrency = flag.Int("concurrency", 10, "The number of threads can be executed concurrently.")
	iteration = flag.Int("iteration", 0, "The number of iterations. If set, it will be terminated within a finite time.")
	identifier = flag.String("id", "anonymous", "An identifier of a client.")
}

func main() {
	var ip string
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}

	flag.Parse()
	fmt.Printf("Start to steal!!! from <%s> by <%s>\n", ip, *identifier)

	limit := limiter.NewConcurrencyLimiter(*concurrency)
	stealer := steal.New()
	steal := func() {
		pubkey, privkey := crypto.GenerateKeyPair()
		addr := crypto.PubkeyToAddress(pubkey)
		limit.Execute(func() { stealer.Steal(addr, privkey) })
	}

	if *iteration > 0 {
		for i := 0; i < *iteration; i++ {
			steal()
		}
		limit.WaitAndClose()
		return
	}

	for {
		steal()
	}
}
