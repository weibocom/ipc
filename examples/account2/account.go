package main

import (
	"log"

	"github.com/kr/pretty"

	"github.com/weibocom/steem-rpc/client"
	"github.com/weibocom/steem-rpc/transports/websocket"
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

	//create(c, "user1")
	//getAccounts(c, "user1")
	//lookupAccountNames(c, "user1")

	lookupAccounts(c, "")
	getAccountCount(c)

	getAccountHistory(c, "user1")
}

func create(c *client.Client, account string) {
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
