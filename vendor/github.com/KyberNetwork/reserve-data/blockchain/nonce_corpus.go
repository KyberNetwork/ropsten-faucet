package blockchain

import (
	ethereum "github.com/ethereum/go-ethereum/common"
	"math/big"
)

type NonceCorpus interface {
	GetAddress() ethereum.Address
	GetNextNonce() (*big.Int, error)
}
