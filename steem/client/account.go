package client

import (
	"github.com/weibocom/ipc/steem/types"
)

// CreateAccount creates a new account.
func (c *Client) CreateAccount(creator string, name string, fee int64, ownerPubKey, activePubKey, postingPubKey, memoPubKey string, jsonMeta string) error {

	operation := &types.AccountCreateOperation{
		Fee:            types.NewSteemAsset(fee),
		Creator:        creator,
		NewAccountName: name,
		Owner:          pubKey2Auth(ownerPubKey),
		Active:         pubKey2Auth(activePubKey),
		Posting:        pubKey2Auth(postingPubKey),
		MemoKey:        types.PublicKey(memoPubKey),
		JsonMetadata:   jsonMeta,
	}

	_, err := c.SendTrx(operation)

	return err
}

func pubKey2Auth(key string) *types.Authority {
	return &types.Authority{
		KeyAuths:        types.KeyAuthorityMap{types.PublicKey(key): 1},
		WeightThreshold: 1,
	}
}
