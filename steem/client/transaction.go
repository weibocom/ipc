package client

import (
	"log"

	"github.com/weibocom/ipc/chain"
	"github.com/weibocom/ipc/steem"
	"github.com/weibocom/ipc/steem/apis/networkbroadcast"
	"github.com/weibocom/ipc/steem/transactions"
	"github.com/weibocom/ipc/steem/types"
)

func (c *Client) CreateTransaction() (*types.Transaction, error) {
	props, err := c.Database.GetDynamicGlobalProperties()
	if err != nil {
		return nil, err
	}
	refBlockPrefix, err := steem.RefBlockPrefix(props.HeadBlockID)
	if err != nil {
		return nil, err
	}

	tx := &types.Transaction{
		RefBlockNum:    steem.RefBlockNum(props.HeadBlockNumber),
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
	tx, err := c.CreateTransaction()
	if err != nil {
		return nil, err
	}
	stx := transactions.NewSignedTransaction(tx)

	for _, op := range operations {
		stx.PushOperation(op)
	}

	if err := stx.Sign(steem.GetPrivateKeys(), chain.MainChain); err != nil {
		log.Printf("transaction sig err:%v\n", err.Error())
		return nil, err
	}

	resp, err = c.NetworkBroadcast.BroadcastTransactionSynchronous(stx.Transaction)
	if err != nil {
		log.Printf("broadcast failed:%v\n", err.Error())
		return nil, err
	}

	return resp, err
}
