package main

import (
	"fmt"
	"os"

	ipcclient "github.com/weibocom/ipc/client"
	"github.com/weibocom/ipc/transports/websocket"
)

func main() {
	tran, err := websocket.NewTransport([]string{"ws://52.80.76.2:8090"})
	if err != nil {
		fmt.Printf("failed to new transport:%s", err.Error())
		os.Exit(-1)
	}

	client, err := ipcclient.NewClient(tran)
	if err != nil {
		fmt.Printf("failed to new client. %v, %v", client, err)
		os.Exit(-2)
	}
	defer client.Close()
	a, err := client.CreateAccount("icy-444", `{"meta":"icy data"}`)
	fmt.Printf("create account:%v, %v", a, err)
}
