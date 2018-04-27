package main

import (
	"fmt"
	"os"

	"github.com/weibocom/ipc/client"
	"github.com/weibocom/ipc/transports/websocket"
)

func main() {
	tran, err := websocket.NewTransport([]string{"ws://52.80.76.2:8090"})
	if err != nil {
		fmt.Println("failed to new transport:%s", err.Error())
		os.Exit(-1)
	}
	defer tran.Close()

	c, cerr := client.NewClient(tran)
	if cerr != nil {
		fmt.Println("failed to new client:%s", cerr.Error())
		os.Exit(-2)
	}
	defer c.Close()

	//testGetActiveWitnesses(c)
}

func testGetWitnesses(c *client.Client) {
	for i := uint32(0); i < 1000000; i++ {
		witnesses, err := c.Condenser.GetWitnesses([]uint32{i})
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%+v\n", witnesses)
	}
}

func testGetWitnessesByAccount(c *client.Client) {
	witness, err := c.Condenser.GetWitnessesByAccount("initminer")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", witness)
}

func testGetActiveWitnesses(c *client.Client) {
	witnesses, err := c.Condenser.GetActiveWitnesses()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", witnesses)
}

func testUpdateWitness(c *client.Client) {

}
