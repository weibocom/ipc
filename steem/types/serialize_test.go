package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weibocom/ipc/encoding"
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

func TestAssetDecode(t *testing.T) {
	// 这个签名是通过 steem 源代码签名获得
	/*
	   asset fee(9, STEEM_SYMBOL)
	   encoder enc;
	   pack(enc, fee);
	   expectedHex = enc.result().str()
	*/
	expectedHex := "3e14d2f1238109c407a634eb31e64371aa905c940c748773630aa6314f958f94"

	fee := NewSteemAsset(9)
	serializedHex, err := encoding.Hash256(fee)
	require.NoError(t, err, "encoding asset ")

	assert.Equal(t, expectedHex, serializedHex, "asset")

}

func TestKeyAuthorityMapDecode(t *testing.T) {
	// 这个签名是通过 steem 源代码签名获得
	/*
	   public_key_type pkt("STM6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF");
	   encoder enc;
	   amap[pkt] = 3;
	   pack(enc, pkt);
	   expectedHex = enc.result().str()
	*/
	expectedHex := "969db862132ac12f3922e8dbc2143386ea2038e0df44bfaeb6f603a887b1d6ee"

	publicKeyStr := "STM6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF"
	kam := NewKeyAuthorityMap()
	kam.AddAuthority(publicKeyStr, 3)

	serializedHex, err := encoding.Hash256(kam)
	require.NoError(t, err, "encoding")

	assert.Equal(t, expectedHex, serializedHex, "key authority map")

}

func TestAuthorityDecode(t *testing.T) {
	// 这个签名是通过 steem 源代码签名获得
	/*
	   auto tx = fc::json::from_file("/data0/steem/programs/util/new_account.json").as<signed_transaction>();
	   //asset fee = asset(9, STEEM_SYMBOL);
	   debug_encoder enc;
	   account_create_operation op = tx.operations[0].get<account_create_operation>();
	   // steem::protocol::authority::key_authority_map amap;
	   // public_key_type pkt("STM6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF");
	   auto owner = op.owner;
	   pack(enc, owner);
	   auto pstr = enc.result().str();
	*/
	expectedHex := "e512b638f11eb8b4bc26a2b5db14ef37cd91f3bb36a35c56c9371d5f6aeaceff"

	ownerPublicKey := "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"

	keyAuths := NewKeyAuthorityMap()
	keyAuths.AddAuthority(ownerPublicKey, 3)

	owner := &Authority{
		KeyAuths:        keyAuths,
		WeightThreshold: 110,
	}

	serializedHex, err := encoding.Hash256(owner)
	require.NoError(t, err, "encoding and hash authority map")

	assert.Equal(t, expectedHex, serializedHex, "encoding hex string")
}

func TestAccountCreateOperationDecode(t *testing.T) {
	/*
	   auto tx = fc::json::from_file("/data0/steem/programs/util/new_account.json").as<signed_transaction>();
	   //asset fee = asset(9, STEEM_SYMBOL);
	   debug_encoder enc;
	   account_create_operation op = tx.operations[0].get<account_create_operation>();
	   // steem::protocol::authority::key_authority_map amap;
	   // public_key_type pkt("STM6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF");
	   // auto owner = op.owner;
	   pack(enc, op);
	   auto pstr = enc.result().str();
	*/
	expectedHex := "d1a44ea478e28d5aea585abdf54b2e7fc5b26c5a6b4bf458640366853ca28da5"
	accountPubKeyStr := "STM6iqZbzYGBnX8mZkn7xK5Z4i7DxcU7GUFo3yWgXuE8BhcbaZpkz"

	newAuthority := func(weight int64, weightThreshold uint32) *Authority {
		keyAuths := NewKeyAuthorityMap()
		keyAuths.AddAuthority(accountPubKeyStr, weight)
		return &Authority{
			KeyAuths:        keyAuths,
			WeightThreshold: weightThreshold,
		}
	}

	op := &AccountCreateOperation{
		Fee:            NewSteemAsset(10),
		Creator:        "initminer",
		NewAccountName: "icy-1",
		Owner:          newAuthority(3, 110),
		Active:         newAuthority(218, 111),
		Posting:        newAuthority(1, 1),
		MemoKey:        PublicKey(accountPubKeyStr),
		JsonMetadata:   `{"meta":"icy data"}`,
	}

	serializedHex, err := encoding.Hash256(op)
	require.NoError(t, err, "encoding and hash transaction")

	assert.Equal(t, expectedHex, serializedHex, "account create operation")
}

func TestTransactionAccountOperationDecode(t *testing.T) {
	/*
	   auto tx = fc::json::from_file("/data0/steem/programs/util/new_account.json").as<transaction>();
	   //asset fee = asset(9, STEEM_SYMBOL);
	   debug_encoder enc;
	   //account_create_operation op = tx.operations[0].get<account_create_operation>();
	   // steem::protocol::authority::key_authority_map amap;
	   // public_key_type pkt("STM6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF");
	   // auto owner = op.owner;
	   pack(enc, tx);
	   auto pstr = enc.result().str();
	*/
	expectedHex := "78343ab5d8702137ba6c7a590cd82dec2655490aa9b1c759c3215a427beab0cd"
	var tx Transaction
	err := json.Unmarshal([]byte(txJson), &tx)
	require.NoError(t, err, "unmarshal transaction json")

	serializedHex, err := encoding.Hash256(tx)
	require.NoError(t, err, "encoding and hash transaction")

	assert.Equal(t, expectedHex, serializedHex, "account create transaction")
}
