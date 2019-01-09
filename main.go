package main

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum/common"
	// raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/cors"
	// "github.com/gin-contrib/sentry"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

type FaucetServer struct {
	r               *gin.Engine
	app             *FaucetApp
	mu              sync.RWMutex
	mu2             sync.RWMutex
	lastVisit       map[string]int64
	session         map[string]*github.User
	sessionIndice   []string
	maxCacheSession int
	oauthConf       *oauth2.Config
}

type SessionManager struct {
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

func (self *FaucetServer) GetUser(c *gin.Context) (*github.User, error) {
	session := sessions.Default(c)
	token := session.Get("token")
	if token == nil {
		return nil, fmt.Errorf("Need to login via github")
	}
	self.mu2.RLock()
	if user, ok := self.session[token.(string)]; ok {
		self.mu2.RUnlock()
		return user, nil
	}
	self.mu2.RUnlock()
	self.mu2.Lock()
	defer self.mu2.Unlock()
	oauthClient := self.oauthConf.Client(oauth2.NoContext, &oauth2.Token{AccessToken: token.(string)})
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		return nil, fmt.Errorf("Could not validate via github. Pls try again")
	}
	if len(self.session) >= self.maxCacheSession {
		oldestSession := self.sessionIndice[0]
		self.sessionIndice = self.sessionIndice[1:]
		delete(self.session, oldestSession)
	}
	self.sessionIndice = append(self.sessionIndice, token.(string))
	self.session[token.(string)] = user
	return user, nil
}

func (self *FaucetServer) Claim(c *gin.Context) {
	if !self.RateGuard(c) {
		return
	}
	user, err := self.GetUser(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"success": false,
				"error":   fmt.Sprintf("Invalid github token"),
			},
		)
		return
	}

	addr := ethereum.HexToAddress(c.PostForm("address"))
	if addr.Big().Cmp(big.NewInt(0)) == 0 {
		c.JSON(
			http.StatusOK,
			gin.H{"success": false, "error": "Invalid address"},
		)
	} else {
		userID := *user.ID
		sent, found := self.app.Get(userID)
		if found {
			c.JSON(
				http.StatusOK,
				gin.H{
					"success": false,
					"error":   fmt.Sprintf("Your account is already registered. We have sent ETH to your account with tx: %s", sent.Hex()),
					"tx":      sent.Hex(),
				},
			)
		} else {
			no, added := self.app.AddAddress(addr, userID)
			if !added {
				latestIndex, yourIndex, _ := self.app.Search(userID)
				c.JSON(
					http.StatusOK,
					gin.H{
						"success": false,
						"error":   fmt.Sprintf("Your account is already registered. If you haven't receive the ETH yet, please wait, there are %d accounts before yours.", yourIndex-latestIndex)},
				)
			} else {
				c.JSON(
					http.StatusOK,
					gin.H{"success": true, "msg": fmt.Sprintf("Your account is added to faucet queue. There are %d accounts before yours. Please wait.", no)},
				)
			}
		}
	}
}

func (self *FaucetServer) SignIn(c *gin.Context) {
	url := self.oauthConf.AuthCodeURL("")
	c.Redirect(http.StatusMovedPermanently, url)
}

func (self *FaucetServer) CallBack(c *gin.Context) {
	code := c.Query("code")
	token, err := self.oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: %s", err)
		return
	}
	session := sessions.Default(c)
	fmt.Println(token.AccessToken)
	session.Set("token", token.AccessToken)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (self *FaucetServer) UserInfo(c *gin.Context) {
	if !self.RateGuard(c) {
		return
	}
	user, err := self.GetUser(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"success": false,
				"error":   err.Error(),
			},
		)
		return
	}
	var (
		msg string
		tx  string
	)
	userID := *user.ID
	sent, found := self.app.Get(userID)
	if found {
		msg = fmt.Sprintf("Your account is already registered. We have sent ETH to your account with tx: %s", sent.Hex())
		tx = sent.Hex()
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"success":    true,
			"user":       user.Login,
			"id":         user.ID,
			"avatar_url": user.AvatarURL,
			"msg":        msg,
			"tx":         tx,
		},
	)
}

func (self *FaucetServer) Run() {
	self.r.GET("/user-info", self.UserInfo)
	self.r.POST("/claim-eth", self.Claim)
	self.r.GET("/sign-in", self.SignIn)
	self.r.GET("/callback", self.CallBack)
	go self.app.Run()
	self.r.Run(":8888")
}

func main() {
	// raven.SetDSN("https://bf15053001464a5195a81bc41b644751:eff41ac715114b20b940010208271b13@sentry.io/228067")
	r := gin.Default()

	// r.Use(sentry.Recovery(raven.DefaultClient, false))
	r.Use(cors.Default())

	secretKey := os.Getenv("SECRET_KEY_BASE")
	store := sessions.NewCookieStore([]byte(secretKey))

	r.Use(sessions.Sessions("faucet", store))

	const maxCacheSession = 300

	server := &FaucetServer{
		r:               r,
		app:             NewFaucetApp(),
		mu:              sync.RWMutex{},
		mu2:             sync.RWMutex{},
		lastVisit:       map[string]int64{},
		session:         make(map[string]*github.User, maxCacheSession),
		sessionIndice:   []string{},
		maxCacheSession: maxCacheSession,
		oauthConf: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			Scopes:       []string{},
			Endpoint:     githuboauth.Endpoint,
		},
	}
	server.Run()
}
