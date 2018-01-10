package main

import (
	"fmt"
	"math/big"
	"net/http"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum/common"
	raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

type FaucetServer struct {
	r         *gin.Engine
	app       *FaucetApp
	mu        sync.Mutex
	lastVisit map[string]int64
}

func (self *FaucetServer) RateGuard(c *gin.Context) bool {
	self.mu.Lock()
	defer self.mu.Unlock()
	ip := c.ClientIP()
	last, found := self.lastVisit[ip]
	cTime := time.Now().Unix()
	if found {
		if cTime-last < 2 {
			// block by limitter
			c.JSON(
				http.StatusOK,
				gin.H{"success": false, "error": "Please try again after 2s, thanks."},
			)
			return false
		}
	}
	self.lastVisit[ip] = cTime
	return true
}

func (self *FaucetServer) Claim(c *gin.Context) {
	if !self.RateGuard(c) {
		return
	}
	addr := ethereum.HexToAddress(c.PostForm("address"))
	if addr.Big().Cmp(big.NewInt(0)) == 0 {
		c.JSON(
			http.StatusOK,
			gin.H{"success": false, "error": "Invalid address"},
		)
	} else {
		sent, found := self.app.Get(addr)
		if found {
			c.JSON(
				http.StatusOK,
				gin.H{
					"success": false,
					"error":   fmt.Sprintf("Your address is already registered. We have sent ETH to your address with tx: %s", sent.Hex()),
					"tx":      sent.Hex(),
				},
			)
		} else {
			no, added := self.app.AddAddress(addr)
			if !added {
				latestIndex, yourIndex, _ := self.app.Search(addr)
				c.JSON(
					http.StatusOK,
					gin.H{
						"success": false,
						"error":   fmt.Sprintf("Your address is already registered. If you haven't receive the ETH yet, please wait, there are %d addresses before yours.", yourIndex-latestIndex)},
				)
			} else {
				c.JSON(
					http.StatusOK,
					gin.H{"success": true, "msg": fmt.Sprintf("Your address is added to faucet queue. There are %d addresses before yours. Please wait.", no)},
				)
			}
		}
	}
}

func (self *FaucetServer) Run() {
	self.r.POST("/claim-eth", self.Claim)
	go self.app.Run()
	self.r.Run(":8888")
}

func main() {
	raven.SetDSN("https://bf15053001464a5195a81bc41b644751:eff41ac715114b20b940010208271b13@sentry.io/228067")
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	r.Use(cors.Default())

	server := &FaucetServer{
		r, NewFaucetApp(), sync.Mutex{}, map[string]int64{},
	}

	server.Run()
}
