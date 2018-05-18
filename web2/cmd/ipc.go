package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/weibocom/ipc/web2/server"
)

var (
	httpAddress = flag.String("http", ":8080", "http address")
	dbAddress   = flag.String("db", "root@/ipc?charset=utf8&parseTime=True&loc=Local&timeout=1s&writeTimeout=3s&readTimeout=3s", "mysql address")
	bcAddress   = flag.String("bc", "ws://52.80.76.2:38090", "blockchain rpc server address")
)

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s := server.New(*httpAddress, *dbAddress, *bcAddress)
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}

	defer s.Stop()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, os.Interrupt)
	<-signalCh

	log.Println("server is closing")
}
