package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum/common"
)

const FILE string = "storage.json"

type pending struct {
	Address ethereum.Address
	User    int64
}

type RamStorage struct {
	mu          sync.RWMutex
	Pending     []*pending
	Registered  map[int64]int
	Sent        map[int64]ethereum.Hash
	LatestIndex int
	Counter     int
}

func NewRamStorage() *RamStorage {
	result := &RamStorage{
		sync.RWMutex{},
		[]*pending{},
		map[int64]int{},
		map[int64]ethereum.Hash{},
		0,
		0,
	}
	go result.RunPersister()
	return result
}

func NewRamStorageFromFile() *RamStorage {
	content, err := ioutil.ReadFile(FILE)
	if err != nil {
		fmt.Printf("Loading database from file failed: %v\n", err)
		return NewRamStorage()
	} else {
		fmt.Printf("Loaded database from %s\n", FILE)
		storage := &RamStorage{
			mu: sync.RWMutex{},
		}
		json.Unmarshal(content, &storage)
		return storage
	}
}

func (self *RamStorage) RunPersister() {
	c := time.Tick(10 * time.Second)
	for _ = range c {
		self.Persist()
	}
}

func (self *RamStorage) Persist() {
	data, err := json.Marshal(self)
	if err == nil {
		err = ioutil.WriteFile(FILE, data, 0644)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		} else {
			fmt.Printf("Persisted to: %s\n", FILE)
		}
	} else {
		fmt.Printf("ERROR: %v\n", err)
	}
}

func (self *RamStorage) Get(user int64) (ethereum.Hash, bool) {
	self.mu.RLock()
	defer self.mu.RUnlock()
	hash, found := self.Sent[user]
	return hash, found
}

func (self *RamStorage) Update(user int64, hash ethereum.Hash) error {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.Sent[user] = hash
	return nil
}

func (self *RamStorage) AddAddress(addr ethereum.Address, user int64) (int, bool) {
	self.mu.Lock()
	defer self.mu.Unlock()
	if _, found := self.Registered[user]; found {
		return 0, false
	} else {
		self.Pending = append(self.Pending, &pending{
			Address: addr,
			User:    user,
		})
		self.Registered[user] = self.LatestIndex
		self.LatestIndex += 1
		return len(self.Pending) - 1, true
	}
}

func (self *RamStorage) GetNextAddress() (ethereum.Address, int64, error) {
	self.mu.Lock()
	defer self.mu.Unlock()
	if len(self.Pending) == 0 {
		return ethereum.Address{}, 0, errors.New("no pending addr left")
	} else {
		result := *self.Pending[0]
		self.Pending[0] = nil
		self.Pending = self.Pending[1:]
		self.Counter = self.Registered[result.User]
		return result.Address, result.User, nil
	}
}

func (self *RamStorage) Search(user int64) (int, int, error) {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.Counter, self.Registered[user], nil
}
