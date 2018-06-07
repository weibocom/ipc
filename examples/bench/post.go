package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/juju/ratelimit"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/satori/go.uuid"
	"github.com/weibocom/ipc/client"
	"github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/store"
	"github.com/weibocom/ipc/transports/websocket"
)

var (
	rpcServer = flag.String("rpc", "ws://10.13.216.128:38090", "steem rpc server")
	n         = flag.Int("n", 2000, "user count")
	initUser  = flag.Bool("user", false, "init users")
	async     = flag.Bool("async", false, "post async or sync")
	batch     = flag.Int("batch", 1, "operation count in one batch")
	cc        = flag.Int("c", runtime.GOMAXPROCS(-1), "concurrency count")
	pr        = flag.Int("pr", 1000, "max rate for post, not used yet")
)

var (
	reg              = metrics.NewRegistry()
	postSuccessMeter = metrics.NewRegisteredTimer("post.success", reg)
	postFailureMeter = metrics.NewRegisteredTimer("post.failure", reg)
)

func main() {
	flag.Parse()

	mysqlStore := store.NewMySQLStore("root@/ipc2?charset=utf8mb4&parseTime=True&loc=Local&timeout=1s&writeTimeout=3s&readTimeout=3s")

	if *initUser {
		tran, err := websocket.NewTransport([]string{*rpcServer},
			websocket.SetAutoReconnectEnabled(true),
			websocket.SetAutoReconnectMaxDelay(1*time.Second),
			websocket.SetReadWriteTimeout(time.Minute),
		)
		if err != nil {
			log.Fatalf("failed to new transport:%s", err.Error())
		}
		defer tran.Close()
		c, cerr := client.NewClient(tran, mysqlStore)
		if cerr != nil {
			log.Fatalf("failed to new client:%s", cerr.Error())
		}

		num := *n / *cc

		var wg sync.WaitGroup
		wg.Add(*cc)

		for i := 0; i < *cc; i++ {
			go initAccounts(&wg, c, i*num, num)
		}

		wg.Wait()

		log.Println("finished to create users")
		return
	}

	go metrics.Log(reg, 10*time.Second, log.New(os.Stdout, "metrics: ", log.Lmicroseconds))
	postLimter := ratelimit.NewBucketWithQuantum(time.Second, int64(*pr), int64(*pr))
	_ = postLimter

	data := []byte("这几天，一封来自北大学生的公开信传播甚广。信中提到的北大相关学院对这位同学提请信息公开一事的处置，引发舆论关注和思考。")

	for i := 0; i < *cc; i++ {
		accPerClient := *n / *cc
		startIndex := i*accPerClient + 1
		go func() {
			tran, err := websocket.NewTransport([]string{*rpcServer})
			if err != nil {
				log.Fatalf("failed to new transport:%s", err.Error())
			}
			defer tran.Close()
			c, cerr := client.NewClient(tran, mysqlStore)
			if cerr != nil {
				log.Fatalf("failed to new client:%s", cerr.Error())
			}

			for j := 0; ; j++ {
				postLimter.Wait(1)
				start := time.Now()

				jj := startIndex + j%accPerClient

				id, _ := uuid.NewV4()
				postid := id.String()

				var dna model.DNA
				if *async {
					_, err = c.PostAsync("wb-"+strconv.Itoa(jj), int64(j), postid, data, postid, []string{"test"})
				} else {
					dna, err = c.Post("wb-"+strconv.Itoa(jj), int64(j), postid, data, postid, []string{"test"})
					// post, err1 := c.LookupPost("wb-"+strconv.Itoa(jj), dna)
					// if err1 != nil || post == nil {
					// 	fmt.Printf("failed to lookup conetent, err : %v\n", err)
					// } else {
					// 	//fmt.Println(post.Content[:3])
					// }
				}

				_ = dna
				if err == nil {
					postSuccessMeter.UpdateSince(start)
				} else {
					postFailureMeter.UpdateSince(start)
					log.Printf("post failed: %v", err)
				}
			}

		}()
	}

	select {}
}

func initAccounts(wg *sync.WaitGroup, c client.Client, start, n int) {
	defer wg.Done()

	fmt.Println("start to init users")
	var err error
	for i := start; i <= start+n; i++ {
		_, err = c.CreateAccount("wb-"+strconv.Itoa(i), `{"meta":"data"}`)
		if err != nil {
			fmt.Printf("create account:%v, %v\n", "wb-"+strconv.Itoa(i), err)
		} else {
			fmt.Printf("create account:%v succeeded\n", "wb-"+strconv.Itoa(i))
		}
	}

	fmt.Println("finished to init users")

}
