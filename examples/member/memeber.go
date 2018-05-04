package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/kr/pretty"

	ipcclient "github.com/weibocom/ipc/client"
	"github.com/weibocom/ipc/store"
	"github.com/weibocom/ipc/transports/websocket"
)

func main() {
	tran, err := websocket.NewTransport([]string{"ws://52.80.76.2:8090"})
	if err != nil {
		log.Fatalf("failed to new transport:%s", err.Error())
	}
	defer tran.Close()

	client, err := ipcclient.NewClient(tran, store.NewMemStore("ipc"))
	if err != nil {
		log.Fatalf("failed to new client. %v, %v", client, err)
	}
	defer client.Close()

	// query
	memebers, err := client.Members()
	if err != nil {
		log.Fatalf("failed to get all memebers: %v", err)
	}

	log.Printf("memebers: %s", pretty.Sprint(memebers))

	// add witness
	rand.Seed(time.Now().UnixNano())
	account := "wbuser-" + strconv.Itoa(rand.Intn(10000))
	m, err := client.AddMember(account)
	if err != nil {
		log.Fatalf("failed to add memeber: %v", err)
	}

	log.Printf("added member:%+v", m)
}
