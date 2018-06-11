package main

// import (
// 	"log"
// 	"math/rand"
// 	"time"

// 	"github.com/kr/pretty"

// 	ipcclient "github.com/weibocom/ipc/client"
// 	"github.com/weibocom/ipc/store"
// 	"github.com/weibocom/ipc/transports/websocket"
// )

// // func init() {
// // 	var defaultConfig = config.IPCConfig{
// // 		Wifs:          []string{"5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX"},
// // 		ChainID:       "2ac122bd70a2f74d6f761c331f4c4da1028b783cc185d23bf5449ac5af49e198",
// // 		AddressPrefix: "STM",
// // 		URL:           "http://weibo.com",
// // 	}

// // 	config.SetConfig(defaultConfig)
// // }
func main() {

	// 	tran, err := websocket.NewTransport([]string{"ws://52.80.76.2:38090"})
	// 	if err != nil {
	// 		log.Fatalf("failed to new transport:%s", err.Error())
	// 	}
	// 	defer tran.Close()

	// 	client, err := ipcclient.NewClient(tran, store.NewMemcacheStore("ipc", "127.0.0.1:11211"))
	// 	if err != nil {
	// 		log.Fatalf("failed to new client. %v, %v", client, err)
	// 	}
	// 	defer client.Close()

	// 	// query
	// 	members, err := client.Members()
	// 	if err != nil {
	// 		log.Fatalf("failed to get all members: %v", err)
	// 	}

	// 	log.Printf("members: %s", pretty.Sprint(members))

	// 	// add witness
	// 	rand.Seed(time.Now().UnixNano())
	// 	//account := "wbuser-" + strconv.Itoa(rand.Intn(10000))
	// 	account := "weibo3"
	// 	m, err := client.AddMember(account)
	// 	if err != nil {
	// 		log.Fatalf("failed to add member: %v", err)
	// 	}

	// 	log.Printf("added member:%+v", m)
}
