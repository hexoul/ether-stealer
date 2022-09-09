package steal

import (
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/hexoul/ether-stealer/v2/infura"
	"github.com/hexoul/ether-stealer/v2/log"
)

var (
	silent *bool
)

func init() {
	silent = flag.Bool("silent", false, "A flag to print failed attempts.")
}

type Stealer struct {
	infuraClient *infura.Infura
	logger       *log.Logger
	silent       bool
}

func New() *Stealer {
	return &Stealer{infuraClient: infura.New(), logger: log.New(), silent: *silent}
}

// Steal steals Ethereum through Infura.
func (s *Stealer) Steal(addr common.Address, privkey []byte) {
	addrStr := addr.String()
	if canStealEther, _ := s.infuraClient.HasBalance(addrStr); canStealEther {
		s.logger.Infof("STEAL Ethereum from the address <%s> with below private key. <%x>", addrStr, privkey)
		// } else if canCandidate, _ := hasTxCount(addrStr); canCandidate {
		// 	log.Infof("GOT CANDIDATE %s !!SECRET!! %x", addrStr, privkey)
		// } else if canStealERC := contract.CanSteal(addr); canStealERC != "" {
		// 	log.Infof("STEAL ERC20 from %s !!SECRET!! %x TARGET %s", addrStr, privkey, canStealERC)
	} else if s.silent == false {
		fmt.Printf("FAILED attempt for address <%s>.\n", addrStr)
	}
}
