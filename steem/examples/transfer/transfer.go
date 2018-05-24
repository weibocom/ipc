package main

import (
	"fmt"
	"log"

	"github.com/weibocom/ipc/steem/client"
	"github.com/weibocom/ipc/transports/websocket"
)

func main() {
	tran, err := websocket.NewTransport([]string{"ws://52.80.76.2:38090"})
	if err != nil {
		log.Fatalf("failed to new transport: %v", err)
	}
	defer tran.Close()

	c, cerr := client.NewClient(tran)
	if cerr != nil {
		log.Fatalf("failed to new client: %v", cerr)
	}
	defer c.Close()

	err = c.Transfer("weibo2", 9000000)

	if err != nil {
		fmt.Printf("failed to transfer to weibo2, err: %v", err)
	} else {
		fmt.Println("transfer succeeded")
	}
}
