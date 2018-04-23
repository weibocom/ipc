package client

import (
	"fmt"
	"log"
	"os"

	"github.com/weibocom/steem-rpc/apis/networkbroadcast"
	"github.com/weibocom/steem-rpc/steem"
	"github.com/weibocom/steem-rpc/transactions"
	"github.com/weibocom/steem-rpc/types"
)

func (c *Client) CreateTransaction() (*types.Transaction, error) {
	props, err := c.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}
	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}

	tx := &types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	}
	return tx, nil
}

func (c *Client) CreateSignedTransaction(creator string, name string, fee int, jsonMeta string) (*transactions.SignedTransaction, error) {
	tx, err := c.CreateTransaction()
	if err != nil {
		return nil, err
	}
	return transactions.NewSignedTransaction(tx), nil
}

// SendTrx signs and sends transactions.
func (c *Client) SendTrx(operations ...types.Operation) (resp *networkbroadcast.BroadcastResponse, err error) {
	props, cfgErr := c.Database.GetDynamicGlobalProperties()
	if cfgErr != nil {
		fmt.Println("failed to get config:%s", cfgErr.Error())
		os.Exit(-3)
	}

	refBlockPrefix, err := transactions.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		fmt.Printf("failed to parse ref block prefix:%v\n", err.Error())
		return
	}

	tx := &types.Transaction{
		RefBlockNum:    transactions.RefBlockNum(props.HeadBlockNumber),
		RefBlockPrefix: refBlockPrefix,
	}

	stx := transactions.NewSignedTransaction(tx)

	for _, op := range operations {
		stx.PushOperation(op)
	}

	if err := stx.Sign(steem.GetPrivateKeys(), steem.SteemChain); err != nil {
		log.Printf("transaction sig err:%v\n", err.Error())
		return nil, err
	}

	resp, err = c.NetworkBroadcast.BroadcastTransactionSynchronous(stx.Transaction)
	if err != nil {
		log.Printf("broadcast failed:%v\n", err.Error())
		return
	}

	return resp, err
}
