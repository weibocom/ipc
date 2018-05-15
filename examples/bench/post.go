package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/juju/ratelimit"
	"github.com/weibocom/ipc/client"
	"github.com/weibocom/ipc/store"
	"github.com/weibocom/ipc/transports/websocket"

	metrics "github.com/rcrowley/go-metrics"
)

var (
	rpcServer = flag.String("rpc", "ws://52.80.76.2:48090", "steem rpc server")
	n         = flag.Int("n", 5000, "user count")
	initUser  = flag.Bool("user", false, "init users")
	batch     = flag.Int("batch", 1, "operation count in one batch")
	cc        = flag.Int("c", runtime.GOMAXPROCS(-1), "concurrency count")
	pr        = flag.Int("pr", 2000, "max rate for post, not used yet")
)

var (
	reg              = metrics.NewRegistry()
	postSuccessMeter = metrics.NewRegisteredTimer("post.success", reg)
	postFailureMeter = metrics.NewRegisteredTimer("post.failure", reg)
)

func main() {
	flag.Parse()
	if *initUser {
		tran, err := websocket.NewTransport([]string{*rpcServer})
		if err != nil {
			log.Fatalf("failed to new transport:%s", err.Error())
		}
		defer tran.Close()
		c, cerr := client.NewClient(tran, store.NewMemStore("weibo"))
		if cerr != nil {
			log.Fatalf("failed to new client:%s", cerr.Error())
		}

		initAccounts(c, *n)
		c.Close()
	}

	go metrics.Log(reg, 10*time.Second, log.New(os.Stdout, "metrics: ", log.Lmicroseconds))
	postLimter := ratelimit.NewBucketWithQuantum(time.Second, int64(*pr), int64(*pr))
	_ = postLimter

	for i := 0; i < *cc; i++ {
		i := i
		go func() {
			tran, err := websocket.NewTransport([]string{*rpcServer})
			if err != nil {
				log.Fatalf("failed to new transport:%s", err.Error())
			}
			defer tran.Close()
			c, cerr := client.NewClient(tran, store.NewMemStore("weibo"))
			if cerr != nil {
				log.Fatalf("failed to new client:%s", cerr.Error())
			}
			defer c.Close()

			for j := 0; ; j++ {
				postLimter.Wait(1)
				start := time.Now()
				jj := j % *n
				_, err := c.Post("wb-"+strconv.Itoa(jj), "人民日报评论：如何聆听“年轻的声音”？​​​"+strconv.Itoa(i)+strconv.Itoa(j),
					[]byte("这几天，一封来自北大学生的公开信传播甚广。信中提到的北大相关学院对这位同学提请信息公开一事的处置，引发舆论关注和思考。"),
					strconv.Itoa(i)+strconv.Itoa(j), []string{"test"})
				if err == nil {
					postSuccessMeter.UpdateSince(start)
				} else {
					postFailureMeter.UpdateSince(start)
				}
			}

		}()
	}

	select {}
}

func initAccounts(c client.Client, n int) {
	fmt.Println("start to init users")
	var err error
	for i := 1; i <= n; i++ {
		_, err = c.CreateAccount("wb-"+strconv.Itoa(i), `{"meta":"data"}`)
		if err != nil {
			fmt.Printf("create account:%v, %v\n", "wb-"+strconv.Itoa(i), err)
		} else {
			fmt.Printf("create account:%v succeeded\n", "wb-"+strconv.Itoa(i))
		}
	}

	fmt.Println("finished to init users")
}
