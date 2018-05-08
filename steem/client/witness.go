package client

import (
	"github.com/weibocom/ipc/steem/types"
)

var maximumBlockSize uint32

func (c *Client) getMaximumBlockSize() (uint32, error) {
	props, err := c.Database.GetDynamicGlobalProperties()
	if err != nil {
		return 0, err
	}

	return uint32(props.MaximumBlockSize.Int64()), nil
}
func (c *Client) AddWitness(privateKeys [][]byte, owner string, blockSigningKey string, url string, fee int64) error {

	if maximumBlockSize == 0 {
		maximumBlockSize, _ = c.getMaximumBlockSize()
	}
	operation := &types.WitnessUpdateOperation{
		Owner:           owner,
		Fee:             types.NewSteemAsset(fee),
		URL:             url,
		BlockSigningKey: types.PublicKey(blockSigningKey),
		Props: &types.ChainProperties{
			AccountCreationFee: types.NewSteemAsset(0),
			MaximumBlockSize:   maximumBlockSize,
			SBDInterestRate:    0,
		},
	}
	_, err := c.SendTrx(privateKeys, operation)

	return err
}
