package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/KyberNetwork/reserve-data/blockchain"
	"github.com/KyberNetwork/reserve-data/blockchain/nonce"
	"github.com/KyberNetwork/reserve-data/signer"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type FaucetApp struct {
	nonce     blockchain.NonceCorpus
	signer    blockchain.Signer
	ethclient *ethclient.Client
	storage   Storage
}

func (self *FaucetApp) Run() {
	c := time.Tick(500 * time.Millisecond)
	for _ = range c {
		addr, err := self.storage.GetNextAddress()
		if err == nil {
			tx, err := self.SendETH(addr)
			fmt.Printf("Tx: %s, err: %v\n", tx.Hex(), err)
		}
	}
}

func (self *FaucetApp) Get(addr ethereum.Address) (ethereum.Hash, bool) {
	return self.storage.Get(addr)
}

func (self *FaucetApp) Update(addr ethereum.Address, hash ethereum.Hash) error {
	return self.storage.Update(addr, hash)
}

func (self *FaucetApp) AddAddress(addr ethereum.Address) (int, bool) {
	return self.storage.AddAddress(addr)
}

func (self *FaucetApp) Search(addr ethereum.Address) (int, int, error) {
	return self.storage.Search(addr)
}

func (self *FaucetApp) SendETH(addr ethereum.Address) (ethereum.Hash, error) {
	option := context.Background()
	amount := big.NewInt(1000000000000000000)
	nonce, err := self.nonce.GetNextNonce()
	if err != nil {
		return ethereum.Hash{}, err
	}
	gasLimit := big.NewInt(50000)
	gasPrice := big.NewInt(30000000000)
	rawTx := types.NewTransaction(
		nonce.Uint64(), addr, amount, gasLimit, gasPrice, []byte{})
	signedTx, err := self.signer.Sign(self.signer.GetAddress(), rawTx)
	if err != nil {
		return ethereum.Hash{}, err
	}
	if err = self.ethclient.SendTransaction(option, signedTx); err != nil {
		return ethereum.Hash{}, err
	}
	self.storage.Update(addr, signedTx.Hash())
	return signedTx.Hash(), nil
}

func NewFaucetApp() *FaucetApp {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fileSigner := signer.NewFileSigner(exPath + "/config.json")
	client, err := rpc.Dial("https://ropsten.infura.io")
	if err != nil {
		panic(err)
	}
	infura := ethclient.NewClient(client)
	nonceCorpus := nonce.NewTimeWindow(infura, fileSigner)
	return &FaucetApp{
		nonceCorpus,
		fileSigner,
		infura,
		NewRamStorageFromFile(),
	}
}
