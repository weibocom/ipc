package types

import (
	// Stdlib
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"

	// RPC
	"github.com/stretchr/testify/assert"
	"github.com/weibocom/ipc/encoding"
)

func TestVoteOperation_Marshal(t *testing.T) {
	op := &VoteOperation{
		Voter:    "xeroc",
		Author:   "xeroc",
		Permlink: "piston",
		Weight:   10000,
	}

	expectedHex := "00057865726f63057865726f6306706973746f6e1027"

	var b bytes.Buffer
	encoder := encoding.NewEncoder(&b)

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
	expectedHex := "cce3032fb2bec232d302d9e242d5e4c129d0d21cb63b65b940ace4b29ba32e81"
	accountPubKeyStr := "6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"

	auth := &Authority{
		KeyAuths:        KeyAuthorityMap{PublicKey(accountPubKeyStr): 1},
		WeightThreshold: 1,
	}

	operation := &AccountCreateOperation{
		Fee:            NewSteemAsset(9),
		Creator:        "initminer",
		NewAccountName: "icy-1",
		Owner:          auth,
		Active:         auth,
		Posting:        auth,
		MemoKey:        PublicKey(accountPubKeyStr),
		JsonMetadata:   `{"meta":"icy data"}`,
	}

	var b bytes.Buffer
	encoder := encoding.NewEncoder(&b)
	encoder.Encode(operation)
	hash := sha256.Sum256(b.Bytes())
	serializedHex := hex.EncodeToString(hash[:])

	assert.Equal(t, expectedHex, serializedHex, "account create operation encoding")
}
