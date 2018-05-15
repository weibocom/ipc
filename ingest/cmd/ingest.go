package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"

	"github.com/weibocom/ipc/ingest"
)

var (
	localIP = ingest.ParseLocalIP()
	cfgFile = flag.String("config", "config.json", "configuration for all url")
	maxTPS  = flag.Int("max-tps", 1000, "max limit of tps")
)

func main() {
	flag.Parse()

	go func() {
		http.ListenAndServe("localhost:9099", nil)
	}()

	buff, err := ioutil.ReadFile(*cfgFile)
	if err != nil {
		fmt.Printf("failed to read config file:%s, err:%v\n", *cfgFile, err)
		return
	}

	config := &ingest.Config{}
	err = json.Unmarshal(buff, config)
	if err != nil {
		fmt.Printf("failed to parse config file:%s, err:%v\n", *cfgFile, err)
		return
	}
	ingest.SetConfig(config)

	done := make(chan struct{}, 1)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		sig := <-sc
		fmt.Printf("%+v ctrl-c(%+v) received. the watcher will exit in seconds.\n", time.Now(), sig)
		close(done)
	}()

	key := fmt.Sprintf(config.Consumer, localIP)
	watcher := ingest.NewConfigWatcher(config.Host, config.Port, key, done, *maxTPS)
	go watcher.Watch()

	<-done
	fmt.Printf("%v done channel closed", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Printf("%+v ingest exit\n", time.Now())
}
