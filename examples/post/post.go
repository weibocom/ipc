package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/kr/pretty"
	"github.com/weibocom/steem-rpc/client"
	"github.com/weibocom/steem-rpc/transports/websocket"
)

func main() {
	tran, err := websocket.NewTransport([]string{"ws://52.80.76.2:8090"})
	if err != nil {
		fmt.Println("failed to new transport:%s", err.Error())
		os.Exit(-1)
	}
	defer tran.Close()

	c, cerr := client.NewClient(tran)
	if cerr != nil {
		fmt.Println("failed to new client:%s", cerr.Error())
		os.Exit(-2)
	}
	defer c.Close()

	// wb-8000 a test tital body body body weibo-8000-9000 wb
	ok, err := c.Post("initminer", "人民日报评论：如何聆听“年轻的声音”？​​​"+strconv.Itoa(rand.Int()), "这几天，一封来自北大学生的公开信传播甚广。信中提到的北大相关学院对这位同学提请信息公开一事的处置，引发舆论关注和思考。", "weibo-8000-9000", "wb", "", []string{"test"})

	//ok, err := c.Post("initminer", "a test post", "a test text", "", "test", "", []string{"test"})

	fmt.Println(ok, err)

	content, err := c.Condenser.GetContent("initminer", "weibo-8000-9000")

	if err != nil {
		log.Fatalf("failed to getConent: %v", err)
	}

	log.Println("content:", pretty.Sprint(content))

	feedEnties, err := c.Condenser.GetFeedEntries("initminer", 0, 1)
	if err != nil {
		log.Fatalf("failed to GetFeedEntries: %v", err)
	}

	log.Println("feedEnties:", pretty.Sprint(feedEnties))

}
