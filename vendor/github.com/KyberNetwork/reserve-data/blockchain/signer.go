package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Signer interface {
	GetAddress() ethereum.Address
	Sign(ethereum.Address, *types.Transaction) (*types.Transaction, error)
	GetTransactOpts() *bind.TransactOpts
}
