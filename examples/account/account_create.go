package main

import (
	"fmt"
	"log"
	"os"

	ipcclient "github.com/weibocom/ipc/client"
	"github.com/weibocom/ipc/store"
	"github.com/weibocom/ipc/transports/websocket"
)

func main() {
	tran, err := websocket.NewTransport([]string{"ws://52.80.76.2:8090"})
	if err != nil {
		fmt.Printf("failed to new transport:%s", err.Error())
		os.Exit(-1)
	}

	client, err := ipcclient.NewClient(tran, store.NewMemStore("ipc"))
	if err != nil {
		fmt.Printf("failed to new client. %v, %v", client, err)
		os.Exit(-2)
	}
	defer client.Close()

	count, err := client.AccountCount()
	if err != nil {
		log.Fatalf("failed to get account count: %v", err)
	}
	log.Printf("current accounts: %d\n", count)

	name := os.Args[1]
	a, err := client.CreateAccount(name, `{"meta":"icy data"}`)
	fmt.Printf("create account:%v, %v\n", a, err)

	a, err = client.CreateAccount(name, `{"meta":"icy data"}`)
	fmt.Printf("create account should failed:%v, %v\n", a, err)
}
