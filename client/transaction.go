package client

import (
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
