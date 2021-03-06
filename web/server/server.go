package server

import (
	"log"
	"math"
	"net"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/keys"
	"github.com/weibocom/ipc/store"
	"github.com/weibocom/ipc/transports/websocket"
	"github.com/weibocom/ipc/web/handler"
	"github.com/weibocom/ipc/web/service"

	ipcclient "github.com/weibocom/ipc/client"
	client "github.com/weibocom/ipc/steem/client"
)

type Server struct {
	ln          net.Listener
	httpAddress string
	dbAddress   string
	bcAddress   string
	company     string

	DB     *gorm.DB
	Client ipcclient.Client
}

func New(httpAddress, dbAddress, bcAddress string, company string) *Server {
	return &Server{
		httpAddress: httpAddress,
		dbAddress:   dbAddress,
		bcAddress:   bcAddress,
		company:     company,
	}
}

func (s *Server) Start() error {
	var err error

	// 1. http
	s.ln, err = net.Listen("tcp", s.httpAddress)
	if err != nil {
		return err
	}
	go http.Serve(s.ln, handler.ConfigRouter())

	// 2. db
	s.DB, err = gorm.Open("mysql", s.dbAddress)
	if err != nil {
		return err
	}

	service.SetDB(s.DB)

	// 3. blockchain
	tran, err := websocket.NewTransport([]string{s.bcAddress}, websocket.SetAutoReconnectEnabled(true), websocket.SetAutoReconnectMaxDelay(time.Minute), websocket.SetReadTimeout(math.MaxInt64))
	if err != nil {
		log.Fatalf("failed to new transport: %v", err)
	}

	chain := client.NewSteemClient(tran, config.GetCreator(), keys.GetPrivateKeys()[0], s.company)
	s.Client, err = ipcclient.NewClient(chain, store.NewMySQLStore(s.dbAddress))
	if err != nil {
		log.Fatalf("failed to new blockchain client: %v", err)
	}
	service.SetIPCClient(s.Client)

	return nil
}

func (s *Server) Stop() {
	err := s.DB.Close()
	if err != nil {
		log.Printf("failed to stop DB: %v", err)
	}

	err = s.Client.Close()
	if err != nil {
		log.Printf("failed to stop blockchain client: %v", err)
	}
}
