package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	ipcclient "github.com/weibocom/ipc/client"
	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/keys"
	"github.com/weibocom/ipc/steem/client"
	"github.com/weibocom/ipc/store"
	"github.com/weibocom/ipc/transports/websocket"
)

func main() {
	tran, err := websocket.NewTransport([]string{"ws://52.80.76.2:48090"})
	if err != nil {
		log.Fatalf("failed to new transport:%s", err.Error())
	}
	defer tran.Close()

	chain := client.NewSteemClient(tran, config.GetCreator(), keys.GetPrivateKeys()[0], "wb")
	client, err := ipcclient.NewClient(chain, store.NewMemStore("ipc"))
	if err != nil {
		log.Fatalf("failed to new client. %v, %v", client, err)
	}
	defer client.Close()

	content := "这几天，一封来自北大学生的公开信传播甚广。信中提到的北大相关学院对这位同学提请信息公开一事的处置，引发舆论关注和思考。"

	rand.Seed(time.Now().UnixNano())
	account := "wbuser-" + strconv.Itoa(rand.Intn(10000))
	a, err := client.CreateAccount(account, `{"meta":"icy data"}`)
	if err != nil {
		log.Fatalf("failed to create account. %v, %v", client, err)
	}

	log.Printf("create account:%v, %v\n", a, err)

	dna, err := client.Post(account, 1000, []byte(content), ipcclient.ContentPost)

	if err != nil {
		log.Fatalf("failed to post: %v", err)
	} else {
		log.Printf("DNA: %s", dna.String())
	}

	content2, err := client.LookupContent(dna)
	if err != nil {
		log.Fatalf("failed to lookup content: %v", err)
	}

	log.Printf("got content: %s", content2)

	existed := client.Verify(dna)
	if !existed {
		log.Fatalln("failed to verify")
	}

	s, err := client.CheckSimilar(dna, dna)
	if err != nil {
		log.Fatalf("failed to CheckSimilar, has one error: %v", err)
	} else if s < 0.9 {
		log.Fatalf("similar of same content should be greater than 0.9, but got %f", s)
	}

	log.Printf("similar: %f", s)
}
