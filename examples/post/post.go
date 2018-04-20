package main

import (
	"fmt"
	"os"

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

	ok, err := c.Post("initminer", "a test post", "a test text", "", "test", "", []string{"test"})

	fmt.Println(ok, err)
}
