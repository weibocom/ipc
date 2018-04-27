package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/kr/pretty"

	"github.com/weibocom/ipc/steem/client"
	"github.com/weibocom/ipc/transports/websocket"
)

func main() {
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
	create(c, account)
	getAccounts(c, account)
	lookupAccountNames(c, account)

	lookupAccounts(c, "")
	getAccountCount(c)

	getAccountHistory(c, account)
}

func create(c *client.Client, account string) {
	ownerPubKey := "STM719WquU4tYem16qxdRrVnzE7F1Xyd8TrGhmcfLFVcG3nvArAdP"
	activePubKey := "STM5bdqA1f1RggUMnBBmcd9vvTMguwWQshkadh8KWuHNX38bSPMif"
	postingPubKey := "STM7zTZCAUvjZZuo9o7FjHsv6irgaAUrWWpR39bWAdpwwpTm8TfbF"
	memoPubKey := "STM6xBkq3xdpvuMKPq6EW2JPXTwoNNDPt6Vqt1w43qCESvBELcniU"

	// ownerPubKey := "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"
	// activePubKey := "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"
	// postingPubKey := "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"
	// memoPubKey := "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"

	err := c.CreateAccount("initminer", account, 1, ownerPubKey, activePubKey, postingPubKey, memoPubKey, `{"meta":"test"}`)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("create account succeeded")
}

func getAccounts(c *client.Client, account string) {
	msg, err := c.Condenser.GetAccounts([]string{account})
	if err != nil {
		log.Fatalf("failed to new client: %v", err)
	}

	log.Println(pretty.Sprint(msg))
}

func lookupAccountNames(c *client.Client, account string) {
	msg, err := c.Condenser.LookupAccountNames([]string{account})
	if err != nil {
		log.Fatalf("failed to new client: %v", err)
	}

	log.Println(pretty.Sprint(msg))
}

func lookupAccounts(c *client.Client, account string) {
	accounts, err := c.Condenser.LookupAccounts(account, 20)
	if err != nil {
		log.Fatalf("failed to new client: %v", err)
	}

	log.Println(accounts)
}

func getAccountCount(c *client.Client) {
	count, err := c.Condenser.GetAccountCount()
	if err != nil {
		log.Fatalf("failed to new client: %v", err)
	}

	log.Println(count)
}

func getAccountHistory(c *client.Client, account string) {
	msg, err := c.Condenser.GetAccountHistory(account, 100, 0)
	if err != nil {
		log.Fatalf("failed to new client: %v", err)
	}

	log.Println(pretty.Sprint(msg))
}
