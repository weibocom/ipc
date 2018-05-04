package client

import (
	"github.com/weibocom/ipc/steem/types"
)

func (c *Client) AddWitness(privateKeys [][]byte, owner string, pubKey string, url string, fee int64) error {
	operation := &types.WitnessUpdateOperation{
		Owner:           owner,
		Fee:             types.NewSteemAsset(fee),
		URL:             url,
		BlockSigningKey: types.PublicKey(pubKey),
		Props: &types.ChainProperties{
			AccountCreationFee: "0 STEEM",
			MaximumBlockSize:   13107200,
			SBDInterestRate:    0,
		},
	}
	_, err := c.SendTrx(privateKeys, operation)

	return err
}
