package transactions

import (
	// Stdlib
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	// RPC
	"github.com/weibocom/steem-rpc/encoding/wif"
	"github.com/weibocom/steem-rpc/steem"
	"github.com/weibocom/steem-rpc/steem/types"
)

// 该测试序列化的 expectedHex 都由steem的源代码计算产生
// 源代码的对象由json反序列化产生。json对应的内容如下：
var txJson = `
{
  "ref_block_num": 12699,
  "ref_block_prefix": 103618507,
  "expiration": "2018-04-12T13:33:22",
  "operations": [
    [
      "account_create",
      {
        "fee": [
          10,
          3,
          "@@000000021"
        ],
        "creator": "initminer",
        "new_account_name": "icy-1",
        "owner": {
          "account_auths": [],
          "key_auths": [
            [
              "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz",
              3
            ]
          ],
          "weight_threshold": 110
        },
        "active": {
          "account_auths": [],
          "key_auths": [
            [
              "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz",
              218
            ]
          ],
          "weight_threshold": 111
        },
        "posting": {
          "account_auths": [],
          "key_auths": [
            [
              "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz",
              1
            ]
          ],
          "weight_threshold": 1
        },
        "memo_key": "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz",
        "json_metadata": "{\"meta\":\"icy data\"}"
      }
    ]
  ],
  "signatures": [
    "1f34e818ea8545a1a55a49cf455b310f722af0d3246448fd63ad383ee9c62e8ebf7b456a4470a1c4de9439348fd0d1512a335d598024e0098e65be1454e02777a3"  ]
}
`

var (
	tx   types.Transaction
	wifs = []string{
		"5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX",
	}
	publicKeys  = make([][]byte, 0, len(wifs))
	privateKeys = make([][]byte, 0, len(wifs))
)

func init() {
	// init transaction
	if err := json.Unmarshal([]byte(txJson), &tx); err != nil {
		panic(err)
	}

	// decode wif
	for _, v := range wifs {
		w, err := wif.DecodeWIF(v)
		if err != nil {
			panic(err)
		}
		privateKeys = append(privateKeys, w.PrivateKey().Serialize())
		publicKeys = append(publicKeys, w.PublicKey().Serialize())
	}

}

func TestTransactionDigest(t *testing.T) {
	expected := "3afe4571043381d00504d4bc4d02ed81f9b0714317edc9cd1c05070b1decdfae"

	stx := NewSignedTransaction(&tx)

	digest, err := stx.Digest(steem.SteemChain)
	if err != nil {
		t.Error(err)
	}

	got := hex.EncodeToString(digest)
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func TestTransactionSignAndVerify(t *testing.T) {
	tx.Signatures = nil
	defer func() {
		tx.Signatures = nil
	}()

	stx := NewSignedTransaction(&tx)
	if err := stx.Sign(privateKeys, steem.SteemChain); err != nil {
		t.Error(err)
	}

	if len(tx.Signatures) != 1 {
		t.Error("expected signatures not appended to the transaction")
	}

	ok, err := stx.Verify(publicKeys, steem.SteemChain)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("verification failed")
	}
	fmt.Printf("sigs:%v\n", stx.Signatures)
}
