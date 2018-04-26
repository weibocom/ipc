package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/kr/pretty"
	"github.com/weibocom/steem-rpc/client"
	"github.com/weibocom/steem-rpc/transports/websocket"
)

func main() {

	// conf := config.SteemConfig{
	// 	Wifs:               []string{"5JNHfZYKGaomSFvd4NUdQ9qMcEAC43kujbfjueTHpVapX1Kzq2n"},
	// 	ChainID:            "18dcf0a285365fc58b71f18b3d3fec954aa0c141c44e4e5cb4cf777b9eab274e",
	// 	SteemAddressPrefix: "TST",
	// }
	// config.SetConfig(conf)

	tran, err := websocket.NewTransport([]string{"ws://52.80.76.2:8090"})
	if err != nil {
		log.Fatalf("failed to new transport: %v", err)
	}
	defer tran.Close()

	c, cerr := client.NewClient(tran)
	if cerr != nil {
		log.Fatalf("failed to new client: %v", cerr)
	}
	defer c.Close()

	rand.Seed(time.Now().UnixNano())
	account := "wbuser-" + strconv.Itoa(rand.Intn(10000))

	log.Println("===========create account============")
	createAccount(c, account)

	log.Println("===========lookup account============")
	lookupAccountNames(c, account)

	permlink := "weibo-8000-9000"

	log.Println("===========add post============")
	addPost(c, account, permlink)

	log.Println("===========get content============")
	getConent(c, account, permlink)

}

func createAccount(c *client.Client, account string) {
	ownerPubKey := "STM719WquU4tYem16qxdRrVnzE7F1Xyd8TrGhmcfLFVcG3nvArAdP"
	activePubKey := "STM5bdqA1f1RggUMnBBmcd9vvTMguwWQshkadh8KWuHNX38bSPMif"
	postingPubKey := "STM7zTZCAUvjZZuo9o7FjHsv6irgaAUrWWpR39bWAdpwwpTm8TfbF"
	memoPubKey := "STM6xBkq3xdpvuMKPq6EW2JPXTwoNNDPt6Vqt1w43qCESvBELcniU"

	_, err := c.CreateAccount("initminer", account, 1, ownerPubKey, activePubKey, postingPubKey, memoPubKey, `{"meta":"test"}`)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("create account succeeded")
}

func lookupAccountNames(c *client.Client, account string) {
	msg, err := c.Condenser.LookupAccountNames([]string{account})
	if err != nil {
		log.Fatalf("failed to lookupAccountNames: %v", err)
	}

	log.Println(pretty.Sprint(msg))
}

func addPost(c *client.Client, account string, permlink string) {
	_, err := c.Post(account, "a test post", "a test text", permlink, "wb", "", []string{"test"})

	if err != nil {
		log.Fatalf("failed to addPost: %v", err)
	}
}

func getConent(c *client.Client, account string, permlink string) {
	content, err := c.Condenser.GetContent(account, permlink)

	if err != nil {
		log.Fatalf("failed to getConent: %v", err)
	}

	log.Println(pretty.Sprint(content))
}
