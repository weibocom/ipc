package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	switcher "git.intra.weibo.com/platform/go-switcher"
	"git.intra.weibo.com/platform/qservice/metrics"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/content"
	"github.com/weibocom/ipc/keys"
	"github.com/weibocom/ipc/web/server"
	"github.com/weibocom/ipc/web/service"
)

var (
	httpAddress    = flag.String("http", ":8080", "http address")
	dbAddress      = flag.String("db", "root@/ipc?charset=utf8mb4&parseTime=True&loc=Local&timeout=1s&writeTimeout=3s&readTimeout=3s", "mysql address")
	bcAddress      = flag.String("bc", "ws://52.80.76.2:38090", "blockchain rpc server address")
	switcherAddr   = flag.String("switcher", "", "switcher addrress")
	graphiteAddr   = flag.String("graphiteAddr", "", "graphite addrress")
	ipcServicePool = flag.String("servicePool", "ipc", "monitor service pool")
	creator        = flag.String("creator", "initminer", "init witness")
	wif            = flag.String("wif", "5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX", "init wif")
)

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	defer content.Clean()

	if *switcherAddr != "" {
		if err := switcher.Serve("tcp", *switcherAddr); err != nil {
			log.Fatalf("serve switch on %s failed: %v", *switcherAddr, err)
		}
	}

	if *graphiteAddr != "" && *ipcServicePool != "" {
		if err := metrics.Init(*graphiteAddr, *ipcServicePool); err != nil {
			log.Fatalf("init metrics module failed: %v", err)
		}
	}

	initConfig()

	s := server.New(*httpAddress, *dbAddress, *bcAddress)
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}

	go service.StartIPCMetrics()

	defer s.Stop()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, os.Interrupt)
	<-signalCh

	log.Println("server is closing")
}

func initConfig() {
	conf := config.GetConfig()
	if *creator != "" {
		conf.Creator = *creator
	}
	if *wif != "" {
		conf.Wifs = []string{*wif}
	}

	config.SetConfig(conf)
	keys.InitKeys(*wif)
}
