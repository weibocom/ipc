package main

import (
	"fmt"
	"os"

	"github.com/weibocom/ipc/chain"
	"github.com/weibocom/ipc/steem"
	"github.com/weibocom/ipc/steem/client"
	"github.com/weibocom/ipc/steem/transactions"
	"github.com/weibocom/ipc/steem/types"
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

	props, cfgErr := c.Database.GetDynamicGlobalProperties()
	if cfgErr != nil {
		fmt.Println("failed to get config:%s", cfgErr.Error())
		os.Exit(-3)
	}

	refBlockPrefix, err := steem.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		fmt.Printf("failed to parse ref block prefix:%v\n", err.Error())
		return
	}

	// expire := time.Date(2018, 4, 14, 17, 8, 35, 0, time.UTC)

	tx := &types.Transaction{
		RefBlockNum:    steem.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
		// RefBlockNum:    29929,
		// RefBlockPrefix: 2639820075,
		// Expiration:     &types.Time{&expire},
	}

	stx := transactions.NewSignedTransaction(tx)

	accountPubKeyStr := "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"

	auth := &types.Authority{
		KeyAuths:        types.KeyAuthorityMap{types.PublicKey(accountPubKeyStr): 1},
		WeightThreshold: 1,
	}

	operation := &types.AccountCreateOperation{
		Fee:            types.NewSteemAsset(9),
		Creator:        "initminer",
		NewAccountName: "icy-1",
		Owner:          auth,
		Active:         auth,
		Posting:        auth,
		MemoKey:        types.PublicKey(accountPubKeyStr),
		JsonMetadata:   `{"meta":"icy data"}`,
	}

	stx.PushOperation(operation)

	if err := stx.Sign(steem.GetPrivateKeys(), chain.MainChain); err != nil {
		fmt.Printf("transaction sig err:%v\n", err.Error())
		return
	}

	// if ok, err := stx.Verify(publicKeys, steem.SteemChain); !ok || err != nil {
	// 	fmt.Printf("verify failed:%v, err:%n", ok, err)
	// }

	resp, err := c.NetworkBroadcast.BroadcastTransactionSynchronous(stx.Transaction)
	if err != nil {
		fmt.Printf("new account failed:%v\n", err.Error())
		return
	}
	fmt.Printf(resp.ID)
}
