package client

import (
	"fmt"

	"github.com/weibocom/steem-rpc/encoding/wif"
	"github.com/weibocom/steem-rpc/types"
)

type Pair struct {
	Public  string
	Private string
}

type Account struct {
	OwnerKey   Pair
	ActiveKey  Pair
	PostingKey Pair
	MemoKey    Pair
}

func (c *Client) NewAccount(creator string, name string, fee int, jsonMeta string) (*Account, error) {
	privKey, err := wif.GenerateKey()
	if err != nil {
		return nil, err
	}
	w, err := wif.NewWIF(privKey)
	if err != nil {
		return nil, err
	}
	wifStr := w.String()
	pubKey := privKey.PublicKey
	fmt.Printf("%v,%v", wifStr, pubKey)
	// accountPubKeyStr := "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"

	// auth := &types.Authority{
	// 	KeyAuths:        types.KeyAuthorityMap{types.PublicKey(accountPubKeyStr): 1},
	// 	WeightThreshold: 1,
	// }

	// operation := &types.AccountCreateOperation{
	// 	Fee:            types.NewSteemAsset(9),
	// 	Creator:        "initminer",
	// 	NewAccountName: "icy-1",
	// 	Owner:          auth,
	// 	Active:         auth,
	// 	Posting:        auth,
	// 	MemoKey:        types.PublicKey(accountPubKeyStr),
	// 	JsonMetadata:   `{"meta":"icy data"}`,
	// }

	// stx.PushOperation(operation)

	// if err := stx.Sign(steem.GetPrivateKeys(), steem.SteemChain); err != nil {
	// 	fmt.Printf("transaction sig err:%v\n", err.Error())
	// 	return
	// }

	// // if ok, err := stx.Verify(publicKeys, steem.SteemChain); !ok || err != nil {
	// //  fmt.Printf("verify failed:%v, err:%n", ok, err)
	// // }

	// resp, err := c.NetworkBroadcast.BroadcastTransactionSynchronous(stx.Transaction)
	// if err != nil {
	// 	fmt.Printf("new account failed:%v\n", err.Error())
	// 	return
	// }
	// fmt.Printf(resp.ID)
	return nil, nil
}

// CreateAccount creates a new account.
func (c *Client) CreateAccount(creator string, name string, fee int64, ownerPubKey, activePubKey, postingPubKey, memoPubKey string, jsonMeta string) (bool, error) {

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

	return err == nil, err
}

func pubKey2Auth(key string) *types.Authority {
	return &types.Authority{
		KeyAuths:        types.KeyAuthorityMap{types.PublicKey(key): 1},
		WeightThreshold: 1,
	}
}
