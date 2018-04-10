package main

import (
	"fmt"
	"os"

	rpc "github.com/icycrystal4/steem-rpc"
	"github.com/icycrystal4/steem-rpc/transactions"
	"github.com/icycrystal4/steem-rpc/transports/websocket"
)

func main() {
	t, err := websocket.NewTransport([]string{"ws://52.80.76.2:8090"})
	if err != nil {
		fmt.Println("failed to new transport:%s", err.Error())
		os.Exit(-1)
	}
	defer t.Close()

	c, cerr := rpc.NewClient(t)
	if cerr != nil {
		fmt.Println("failed to new client:%s", cerr.Error())
		os.Exit(-2)
	}
	defer c.Close()

	obj, cfgErr := c.Database.GetDynamicGlobalProperties()
	if cfgErr != nil {
		fmt.Println("failed to get config:%s", cfgErr.Error())
		os.Exit(-3)
	}

	fmt.Print("block head:", obj.HeadBlockID)

	var tx *transactions.SignedTransaction

}
