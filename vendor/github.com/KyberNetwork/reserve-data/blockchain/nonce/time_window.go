package nonce

import (
	"context"
	"math/big"
	"sync"

	"github.com/KyberNetwork/reserve-data/blockchain"
	"github.com/KyberNetwork/reserve-data/common"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TimeWindow struct {
	ethclient   *ethclient.Client
	signer      blockchain.Signer
	mu          sync.Mutex
	manualNonce *big.Int
	time        uint64 `last time nonce was requested`
}

func NewTimeWindow(
	ethclient *ethclient.Client,
	signer blockchain.Signer) *TimeWindow {
	return &TimeWindow{
		ethclient,
		signer,
		sync.Mutex{},
		big.NewInt(0),
		0,
	}
}

func (self *TimeWindow) GetAddress() ethereum.Address {
	return self.signer.GetAddress()
}

func (self *TimeWindow) getNonceFromNode() (*big.Int, error) {
	option := context.Background()
	nonce, err := self.ethclient.PendingNonceAt(option, self.signer.GetAddress())
	return big.NewInt(int64(nonce)), err
}

func (self *TimeWindow) GetNextNonce() (*big.Int, error) {
	self.mu.Lock()
	defer self.mu.Unlock()
	t := common.GetTimepoint()
	if t-self.time < 2000 {
		self.time = t
		self.manualNonce.Add(self.manualNonce, ethereum.Big1)
		return self.manualNonce, nil
	} else {
		nonce, err := self.getNonceFromNode()
		if err != nil {
			return big.NewInt(0), err
		}
		self.time = t
		self.manualNonce = nonce
		return self.manualNonce, nil
	}
}
