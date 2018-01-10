package main

import (
	ethereum "github.com/ethereum/go-ethereum/common"
)

type Storage interface {
	AddAddress(addr ethereum.Address) (int, bool)
	GetNextAddress() (ethereum.Address, error)
	Search(addr ethereum.Address) (int, int, error)
	Get(addr ethereum.Address) (ethereum.Hash, bool)
	Update(addr ethereum.Address, hash ethereum.Hash) error
}
