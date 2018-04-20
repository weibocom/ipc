package wif

import (
	// Stdlib

	"encoding/hex"
	mr "math/rand"
	"testing"
	"time"

	"github.com/weibocom/steem-rpc/encoding/wif"
)

type testData struct {
	WIF           string
	PrivateKeyHex string
}

var TestData = []testData{
	{
		WIF:           "5JWHY5DxTF6qN5grTtChDCYBmWHfY9zaSsw4CxEKN5eZpH9iBma",
		PrivateKeyHex: "5ad2b8df2c255d4a2996ee7d065e013e1bbb35c075ee6e5208aca44adc9a9d4c",
	},
	{
		WIF:           "5KPipdRzoxrp6dDqsBfMD6oFZG356trVHV5QBGx3rABs1zzWWs8",
		PrivateKeyHex: "cf9d6121ed458f24ea456ad7ff700da39e86688988cfe5c6ed6558642cf1e32f",
	},
}

func init() {
	mr.Seed(time.Now().UnixNano())
}

func TestDecode(t *testing.T) {
	for _, d := range TestData {
		privKey, err := wif.Decode(d.WIF)
		if err != nil {
			t.Error(err)
		}

		expected := d.PrivateKeyHex
		got := hex.EncodeToString(privKey)

		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	}
}

func TestPublicKey(t *testing.T) {
	wifStr := "5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX"
	// 不带 STM 等前缀
	pubkeyStr := "6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF"

	w, err := wif.DecodeWIF(wifStr)
	if err != nil {
		t.Error(err)
		return
	}
	var pubKey wif.PublicKey
	err = pubKey.From(pubkeyStr)
	if err != nil {
		t.Error(err)
		return
	}

	pubks := pubKey.String(false)
	if pubks != pubkeyStr {
		t.Errorf("expect %s but got %s", pubkeyStr, pubks)
	}

	pubksFromWif := w.PublicKey().String(false)
	if pubksFromWif != pubkeyStr {
		t.Errorf("expect %s but got %s", pubkeyStr, pubksFromWif)
	}

}

func TestParser(t *testing.T) {
	wifStr := "5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX"
	// pubkeyStr := "6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF"
	w, err := wif.DecodeWIF(wifStr)
	if err != nil {
		t.Error(err)
		return
	}
	wstr := w.String()
	if wstr != wifStr {
		t.Errorf("expect %s but got %s", wifStr, wstr)
		return
	}
	priv, err := wif.GenerateKey()
	if err != nil {
		t.Error(err)
		return
	}
	_, err = wif.NewWIF(priv)
	if err != nil {
		t.Error(err)
	}

}

func TestNewWIF(t *testing.T) {
	pk, err := wif.GenerateKey()
	if err != nil {
		t.Error(err)
		return
	}
	_, err = wif.NewWIF(pk)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestWifParser(t *testing.T) {
	wifStr := "5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX"
	pubkeyStr := "STM6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF"
	w, err := wif.DecodeWIF(wifStr)
	if err != nil {
		t.Error(err)
		return
	}
	recoverPubKeyStr := w.PublicKey().String(true)

	if pubkeyStr != recoverPubKeyStr {
		t.Errorf("expected %v but got %v", pubkeyStr, recoverPubKeyStr)
	}
}
