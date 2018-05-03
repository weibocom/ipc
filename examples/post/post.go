package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

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

	title := "人民日报评论：如何聆听“年轻的声音"
	content := "这几天，一封来自北大学生的公开信传播甚广。信中提到的北大相关学院对这位同学提请信息公开一事的处置，引发舆论关注和思考。"
	uri := "http://weibo.com"

	rand.Seed(time.Now().UnixNano())
	account := "wbuser-" + strconv.Itoa(rand.Intn(10000))
	a, err := client.CreateAccount(account, `{"meta":"icy data"}`)
	if err != nil {
		log.Fatalf("failed to create account. %v, %v", client, err)
	}

	log.Printf("create account:%v, %v\n", a, err)

	dna, err := client.Post(account, title, []byte(content), uri, []string{"人民日报", "北大学生", "公开信", "年轻的声音"})

	if err != nil {
		log.Fatalf("failed to post: %v", err)
	} else {
		log.Printf("DNA: %s", dna.ID())
	}
}
