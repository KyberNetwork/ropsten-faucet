package blockchain

import (
	"context"
	"fmt"
	"log"

	"github.com/KyberNetwork/reserve-data/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Blockchain struct {
	client  *ethclient.Client
	wrapper *ContractWrapper
	reserve *ReserveContract
	rm      ethereum.Address
	signer  Signer
	tokens  []common.Token
	nonce   NonceCorpus
}

func (self *Blockchain) AddToken(t common.Token) {
	self.tokens = append(self.tokens, t)
}

func (self *Blockchain) IsMined(tx ethereum.Hash) (bool, error) {
	option := context.Background()
	receipt, err := self.client.TransactionReceipt(option, tx)
	return receipt != nil, err
}

func (self *Blockchain) getTransactOpts() (*bind.TransactOpts, error) {
	shared := self.signer.GetTransactOpts()
	nonce, err := self.nonce.GetNextNonce()
	if err != nil {
		return nil, err
	} else {
		result := bind.TransactOpts{
			shared.From,
			nonce,
			shared.Signer,
			shared.Value,
			shared.GasPrice,
			shared.GasLimit,
			shared.Context,
		}
		return &result, nil
	}
}

func (self *Blockchain) FetchBalanceData(reserve ethereum.Address, timepoint uint64) (map[string]common.BalanceEntry, error) {
	result := map[string]common.BalanceEntry{}
	tokens := []ethereum.Address{}
	for _, tok := range self.tokens {
		tokens = append(tokens, ethereum.HexToAddress(tok.Address))
	}
	timestamp := common.GetTimestamp()
	balances, err := self.wrapper.GetBalances(nil, reserve, tokens)
	returnTime := common.GetTimestamp()
	log.Printf("Fetcher ------> balances: %v, err: %s", balances, err)
	if err != nil {
		for tokenID, _ := range common.SupportedTokens {
			result[tokenID] = common.BalanceEntry{
				Valid:      false,
				Error:      err.Error(),
				Timestamp:  timestamp,
				ReturnTime: returnTime,
			}
		}
	} else {
		for i, tok := range self.tokens {
			result[tok.ID] = common.BalanceEntry{
				Valid:      true,
				Timestamp:  timestamp,
				ReturnTime: returnTime,
				Balance:    common.RawBalance(*balances[i]),
			}
		}
	}
	return result, nil
}

func (self *Blockchain) FetchRates(
	sources []common.Token,
	dests []common.Token, timepoint uint64) (common.AllRateEntry, error) {

	result := common.AllRateEntry{}
	sourceAddrs := []ethereum.Address{}
	for _, s := range sources {
		sourceAddrs = append(sourceAddrs, ethereum.HexToAddress(s.Address))
	}
	destAddrs := []ethereum.Address{}
	for _, d := range dests {
		destAddrs = append(destAddrs, ethereum.HexToAddress(d.Address))
	}
	timestamp := common.GetTimestamp()
	rates, expiries, balances, err := self.wrapper.GetPrices(
		nil, self.rm, sourceAddrs, destAddrs)
	// fmt.Printf("\nrates (%d): %v\n", len(rates), rates)
	// fmt.Printf("expiries: %v\n", expiries)
	// fmt.Printf("balances: %v\n", balances)
	// fmt.Printf("error: %s\n", err)
	returnTime := common.GetTimestamp()
	result.Timestamp = timestamp
	result.ReturnTime = returnTime
	if err != nil {
		result.Valid = false
		result.Error = err.Error()
		return result, err
	} else {
		result.Valid = true
		result.Data = map[common.TokenPairID]common.RateEntry{}
		for i, s := range sources {
			result.Data[common.NewTokenPairID(
				s.ID, dests[i].ID)] = common.RateEntry{
				rates[i], expiries[i], balances[i],
				// rates[3*i], rates[3*i+1], rates[3*i+2],
			}
		}
		return result, nil
	}
}

func (self *Blockchain) SetRates(
	sources []ethereum.Address,
	dests []ethereum.Address,
	rates []*big.Int,
	expiryBlocks []*big.Int) (ethereum.Hash, error) {

	opts, err := self.getTransactOpts()
	if err != nil {
		fmt.Printf("Getting transaction opts failed!!!!!!!\n")
		return ethereum.Hash{}, err
	} else {
		tx, err := self.reserve.SetRate(
			opts,
			sources, dests, rates, expiryBlocks, true)
		if err != nil {
			fmt.Printf("Broadcasting transaction failed!!!!!!!\n")
			return ethereum.Hash{}, err
		} else {
			return tx.Hash(), err
		}
	}
}

func (self *Blockchain) Send(
	token common.Token,
	amount *big.Int,
	dest ethereum.Address) (ethereum.Hash, error) {

	opts, err := self.getTransactOpts()
	if err != nil {
		return ethereum.Hash{}, err
	} else {
		tx, err := self.reserve.Withdraw(
			opts,
			ethereum.HexToAddress(token.Address),
			amount, dest)
		if err != nil {
			return ethereum.Hash{}, err
		} else {
			return tx.Hash(), err
		}
	}
}

// func (self *Blockchain) sendToken(token common.Token, amount *big.Int, address ethereum.Address) (ethereum.Hash, error) {
// 	erc20, err := NewErc20Contract(
// 		ethereum.HexToAddress(token.Address),
// 		self.ethclient,
// 	)
// 	fmt.Printf("address: %s\n", token.Address)
// 	if err != nil {
// 		return ethereum.Hash{}, err
// 	}
// 	tx, err := erc20.Transfer(
// 		self.signer.GetTransactOpts(),
// 		address, amount)
// 	if err != nil {
// 		return ethereum.Hash{}, err
// 	} else {
// 		return tx.Hash(), nil
// 	}
// }
//
// func (self *Blockchain) sendETH(
// 	amount *big.Int,
// 	address ethereum.Address) (ethereum.Hash, error) {
// 	// nonce, gasLimit, gasPrice gets from ethclient
//
// 	option := context.Background()
// 	rm := self.signer.GetAddress()
// 	nonce, err := self.ethclient.PendingNonceAt(
// 		option, rm)
// 	if err != nil {
// 		return ethereum.Hash{}, err
// 	}
// 	gasLimit := big.NewInt(1000000)
// 	gasPrice := big.NewInt(20000000000)
// 	rawTx := types.NewTransaction(
// 		nonce, address, amount, gasLimit, gasPrice, []byte{})
// 	signedTx, err := self.signer.Sign(rm, rawTx)
// 	if err != nil {
// 		return ethereum.Hash{}, err
// 	}
// 	if err = self.ethclient.SendTransaction(option, signedTx); err != nil {
// 		return ethereum.Hash{}, err
// 	}
// 	return signedTx.Hash(), nil
// }

func NewBlockchain(
	ethereum *ethclient.Client,
	wrapperAddr, reserveAddr ethereum.Address,
	signer Signer, nonceCorpus NonceCorpus) (*Blockchain, error) {
	wrapper, err := NewContractWrapper(wrapperAddr, ethereum)
	if err != nil {
		return nil, err
	}
	log.Printf("reserve owner address: %s", signer.GetAddress().Hex())
	log.Printf("reserve address: %s", reserveAddr.Hex())
	reserve, err := NewReserveContract(reserveAddr, ethereum)
	if err != nil {
		return nil, err
	}
	return &Blockchain{
		client:  ethereum,
		wrapper: wrapper,
		reserve: reserve,
		rm:      reserveAddr,
		signer:  signer,
		tokens:  []common.Token{},
		nonce:   nonceCorpus,
	}, nil
}
