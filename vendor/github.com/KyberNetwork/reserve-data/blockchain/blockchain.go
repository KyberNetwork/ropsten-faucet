package blockchain

import (
	"context"
	"log"
	"strconv"

	"github.com/KyberNetwork/reserve-data/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type tbindex struct {
	BulkIndex   uint64
	IndexInBulk uint64
}

type Blockchain struct {
	rpcClient    *rpc.Client
	client       *ethclient.Client
	wrapper      *ContractWrapper
	pricing      *Pricing
	reserve      *ReserveContract
	rm           ethereum.Address
	pricingAddr  ethereum.Address
	signer       Signer
	tokens       []common.Token
	tokenIndices map[string]tbindex
	nonce        NonceCorpus
}

func (self *Blockchain) AddToken(t common.Token) {
	self.tokens = append(self.tokens, t)
}

func (self *Blockchain) LoadAndSetTokenIndices() error {
	tokens := []ethereum.Address{}
	for _, tok := range self.tokens {
		tokens = append(tokens, ethereum.HexToAddress(tok.Address))
	}
	bulkIndices, indicesInBulk, err := self.wrapper.GetTokenIndicies(
		nil,
		self.pricingAddr,
		tokens,
	)
	if err != nil {
		return err
	}
	self.tokenIndices = map[string]tbindex{}

	for i, tok := range tokens {
		self.tokenIndices[tok.Hex()] = tbindex{
			bulkIndices[i].Uint64(),
			indicesInBulk[i].Uint64(),
		}
	}
	log.Printf("Token indices: %+v", self.tokenIndices)
	return nil
}

func (self *Blockchain) CurrentBlock() (uint64, error) {
	var blockno string
	err := self.rpcClient.Call(&blockno, "eth_blockNumber")
	if err != nil {
		return 0, err
	}
	result, err := strconv.ParseUint(blockno, 0, 64)
	return result, err
}

func (self *Blockchain) IsMined(tx ethereum.Hash) (bool, error) {
	option := context.Background()
	receipt, err := self.client.TransactionReceipt(option, tx)
	if receipt != nil {
		err = nil
	}
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

func (self *Blockchain) FetchRates(timepoint uint64) (common.AllRateEntry, error) {
	result := common.AllRateEntry{}
	tokenAddrs := []ethereum.Address{}
	for _, s := range self.tokens {
		tokenAddrs = append(tokenAddrs, ethereum.HexToAddress(s.Address))
	}
	timestamp := common.GetTimestamp()
	baseBuys, baseSells, compactBuys, compactSells, blocks, err := self.wrapper.GetTokenRates(
		nil, self.pricingAddr, tokenAddrs,
	)
	returnTime := common.GetTimestamp()
	result.Timestamp = timestamp
	result.ReturnTime = returnTime
	if err != nil {
		result.Valid = false
		result.Error = err.Error()
		return result, err
	} else {
		result.Valid = true
		result.Data = map[string]common.RateEntry{}
		for i, token := range self.tokens {
			result.Data[token.ID] = common.RateEntry{
				baseBuys[i],
				int8(compactBuys[i]),
				baseSells[i],
				int8(compactSells[i]),
				blocks[i].Uint64(),
			}
		}
		return result, nil
	}
}

func (self *Blockchain) GetPrice(token ethereum.Address, block *big.Int, priceType string, qty *big.Int) (*big.Int, error) {
	if priceType == "buy" {
		return self.pricing.GetPrice(nil, token, block, true, qty)
	} else {
		return self.pricing.GetPrice(nil, token, block, false, qty)
	}
}

func (self *Blockchain) SetRates(
	tokens []ethereum.Address,
	buys []*big.Int,
	sells []*big.Int,
	block *big.Int) (ethereum.Hash, error) {

	opts, err := self.getTransactOpts()
	block.Add(block, big.NewInt(1))
	if err != nil {
		log.Printf("Getting transaction opts failed!!!!!!!\n")
		return ethereum.Hash{}, err
	} else {
		baseBuys, baseSells, compactBuys, compactSells, _, err := self.wrapper.GetTokenRates(
			nil, self.pricingAddr, tokens,
		)
		if err != nil {
			return ethereum.Hash{}, err
		}
		baseTokens := []ethereum.Address{}
		newBSells := []*big.Int{}
		newBBuys := []*big.Int{}
		newCSells := map[ethereum.Address]byte{}
		newCBuys := map[ethereum.Address]byte{}
		for i, token := range tokens {
			compactSell, overflow1 := BigIntToCompactRate(sells[i], baseSells[i])
			compactBuy, overflow2 := BigIntToCompactRate(buys[i], baseBuys[i])
			if overflow1 || overflow2 {
				baseTokens = append(baseTokens, token)
				newBSells = append(newBSells, compactSell.Base)
				newBBuys = append(newBBuys, compactBuy.Base)
			} else {
				if compactSell.Compact != byte(compactSells[i]) ||
					compactBuy.Compact != byte(compactBuys[i]) {
					newCSells[token] = compactSell.Compact
					newCBuys[token] = compactBuy.Compact
				}
			}
		}
		buys, sells, indices := BuildCompactBulk(
			newCBuys,
			newCSells,
			self.tokenIndices,
		)
		var tx *types.Transaction
		if len(baseTokens) > 0 {
			// set base tx
			tx, err = self.pricing.SetBasePrice(
				opts, baseTokens, newBBuys, newBSells,
				buys, sells, block, indices)
			log.Printf("Setting base rates: tx(%s), err(%v) with baseTokens(%+v), basebuys(%+v), basesells(%+v), buys(%+v), sells(%+v), block(%s), indices(%+v)",
				tx.Hash().Hex(), err, baseTokens, newBBuys, newBSells, buys, sells, block.Text(10), indices,
			)
		} else {
			// update compact tx
			tx, err = self.pricing.SetCompactData(
				opts, buys, sells, block, indices)
			log.Printf("Setting compact rates: tx(%s), err(%v) with basesells(%+v), buys(%+v), sells(%+v), block(%s), indices(%+v)",
				tx.Hash().Hex(), err, baseTokens, buys, sells, block.Text(10), indices,
			)
		}
		if err != nil {
			log.Printf("Broadcasting transaction failed!!!!!!!\n")
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
	client *rpc.Client,
	ethereum *ethclient.Client,
	wrapperAddr, pricingAddr, reserveAddr ethereum.Address,
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
	log.Printf("pricing address: %s", pricingAddr.Hex())
	pricing, err := NewPricing(pricingAddr, ethereum)
	if err != nil {
		return nil, err
	}
	return &Blockchain{
		rpcClient:   client,
		client:      ethereum,
		wrapper:     wrapper,
		pricing:     pricing,
		reserve:     reserve,
		rm:          reserveAddr,
		pricingAddr: pricingAddr,
		signer:      signer,
		tokens:      []common.Token{},
		nonce:       nonceCorpus,
	}, nil
}
