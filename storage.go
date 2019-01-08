package main

import (
	ethereum "github.com/ethereum/go-ethereum/common"
)

type Storage interface {
	AddAddress(addr ethereum.Address, user int64) (int, bool)
	GetNextAddress() (ethereum.Address, int64, error)
	Search(user int64) (int, int, error)
	Get(user int64) (ethereum.Hash, bool)
	Update(user int64, hash ethereum.Hash) error
}
