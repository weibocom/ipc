package types

import (
	// Stdlib
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	// RPC
	"github.com/weibocom/ipc/encoding/transaction"
	"github.com/weibocom/ipc/steem/types"

	base58 "github.com/itchyny/base58-go"
)

func TestVoteOperation_MarshalTransaction(t *testing.T) {
	op := &types.VoteOperation{
		Voter:    "xeroc",
		Author:   "xeroc",
		Permlink: "piston",
		Weight:   10000,
	}

	expectedHex := "00057865726f63057865726f6306706973746f6e1027"

	var b bytes.Buffer
	encoder := transaction.NewEncoder(&b)

	if err := encoder.Encode(op); err != nil {
		t.Error(err)
	}

	serializedHex := hex.EncodeToString(b.Bytes())

	if serializedHex != expectedHex {
		t.Errorf("expected %v, got %v", expectedHex, serializedHex)
	}
}

func TestAccountCreateOperation(t *testing.T) {
	// 这个签名是通过steem的pack与encoder进行序列化，hash之后的值
	// expectedHex := "cce3032fb2bec232d302d9e242d5e4c129d0d21cb63b65b940ace4b29ba32e81"
	accountPubKeyStr := "6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"

	auth := &types.Authority{
		KeyAuths:        types.KeyAuthorityMap{accountPubKeyStr: 1},
		WeightThreshold: 1,
	}

	operation := &types.AccountCreateOperation{
		Fee:            types.NewSteemAsset(9),
		Creator:        "initminer",
		NewAccountName: "icy-1",
		Owner:          auth,
		Active:         auth,
		Posting:        auth,
		MemoKey:        accountPubKeyStr,
		JsonMetadata:   `{"meta":"icy data"}`,
	}

	// var b bytes.Buffer
	// encoder := transaction.NewEncoder(&b)
	// encoder.Encode(operation.Owner)
	// hash := sha256.Sum256(b.Bytes())
	// serializedHex := hex.EncodeToString(hash[:])

	// if serializedHex != expectedHex {
	// 	t.Errorf("expected %v, got %v", expectedHex, serializedHex)
	// }
	// fmt.Printf("bytes:\n%v\n", b.Bytes())

	encoding := base58.FlickrEncoding
	d, e := encoding.Decode([]byte(accountPubKeyStr))
	fmt.Printf("decode public key:%d, err:%v\n", len(d), e)
}
