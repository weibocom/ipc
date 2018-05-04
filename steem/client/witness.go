package client

import (
	"github.com/weibocom/ipc/steem/types"
)

func (c *Client) AddWitness(privateKeys [][]byte, owner string, blockSigningKey string, url string, fee int64) error {
	operation := &types.WitnessUpdateOperation{
		Owner:           owner,
		Fee:             types.NewSteemAsset(fee),
		URL:             url,
		BlockSigningKey: types.PublicKey(blockSigningKey),
		Props: &types.ChainProperties{
			AccountCreationFee: types.NewSteemAsset(0),
			MaximumBlockSize:   13107200,
			SBDInterestRate:    0,
		},
	}
	_, err := c.SendTrx(privateKeys, operation)

	return err
}
