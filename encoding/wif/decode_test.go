package wif

import (
	// Stdlib
	"bytes"
	"encoding/hex"
	"fmt"
	mr "math/rand"
	"testing"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/weibocom/steem-rpc/encoding/wif"
)

func init() {
	mr.Seed(time.Now().UnixNano())
}

func TestDecode(t *testing.T) {
	for _, d := range wif.TestData {
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
	direct, l := wif.ParsePubKeyBase58(pubkeyStr)

	fromPriKey, err := wif.GetPublicKey(wifStr)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(fromPriKey, direct) {
		t.Errorf("public key parse failed. \n\tpublic key generated from private key:%v\n\tpublic key directly parse from base58:%v\n", fromPriKey, direct)
	}
	if l != len(direct) {
		t.Logf("public key parse from base58 maybe trunced from %d to %d", l, len(direct))
	}
}

func TestParser(t *testing.T) {
	wifStr := "5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX"
	// pubkeyStr := "6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF"
	w, err := wif.ParseWif(wifStr)
	if err != nil {
		t.Error(err)
		return
	}
	pk := w.PrivKey.Serialize()
	fmt.Printf("private key:%v len:%d\n", pk, len(pk))
	fmt.Printf("wif base 58:%v\n", base58.Decode(wifStr))

	priv, err := wif.NewPrivateKey()
	if err != nil {
		t.Error(err)
		return
	}
	wif1, err := btcutil.NewWIF(priv, &chaincfg.MainNetParams, false)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("witness:%v\n", wif1.String())
}
