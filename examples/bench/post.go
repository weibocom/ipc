package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/weibocom/ipc/keys"

	"github.com/juju/ratelimit"
	"github.com/weibocom/ipc/steem/client"
	"github.com/weibocom/ipc/steem/types"
	"github.com/weibocom/ipc/transports/websocket"

	metrics "github.com/rcrowley/go-metrics"
)

var (
	rpcServer = flag.String("rpc", "ws://52.80.76.2:8090", "steem rpc server")
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
			c, cerr := client.NewClient(tran)
			if cerr != nil {
				log.Fatalf("failed to new client:%s", cerr.Error())
				os.Exit(-2)
			}
			defer c.Close()

			for j := 0; ; j++ {
				postLimter.Wait(1)
				start := time.Now()

				var ops []types.Operation
				for k := 0; k < *batch; k++ {
					op := client.CreateCommentOperation("initminer", "人民日报评论：如何聆听“年轻的声音”？​​​"+strconv.Itoa(i)+strconv.Itoa(j)+strconv.Itoa(k), "这几天，一封来自北大学生的公开信传播甚广。信中提到的北大相关学院对这位同学提请信息公开一事的处置，引发舆论关注和思考。", "weibo-8000-9000", "wb", "", []string{"test"})
					ops = append(ops, op)

				}
				_, err = c.BatchPost(keys.GetPrivateKeys(), ops)

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
