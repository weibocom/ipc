package main

import (
	"fmt"
	"log"
	"os"

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
	ok, err := c.Post("initminer", "a test post", "a test text", "weibo-8000-9000", "wb", "", []string{"test"})

	//ok, err := c.Post("initminer", "a test post", "a test text", "", "test", "", []string{"test"})

	fmt.Println(ok, err)

	content, err := c.Condenser.GetContent("initminer", "weibo-8000-9000")

	if err != nil {
		log.Fatalf("failed to getConent: %v", err)
	}

	log.Println(pretty.Sprint(content))
}
