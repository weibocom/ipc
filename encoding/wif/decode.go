package wif

import (
	// Vendor

	"crypto/ecdsa"
	"crypto/rand"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"
)

// Decode can be used to turn WIF into a raw private key (32 bytes).
func Decode(wif string) ([]byte, error) {
	w, err := DecodeWIF(wif)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode WIF")
	}

	return w.Serialize(), nil
}

func DecodeWIF(wif string) (*WIF, error) {
	w, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return nil, err
	}
	return (*WIF)(w), nil
}

func ParsePrivateKey(pk []byte) *btcec.PrivateKey {
	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), pk)
	return privateKey
}

// ParsePublicKey returns the public key associated with the given public-base58-formatted stringent key
// in the 33-byte compressed format.
// pubkeyStr  不包含任何前缀。比如STM等。
func ParsePubKeyBase58(pubkeyStr string) ([]byte, int) {
	// TODO: 该函数返回的字节数并不是33，而是37. 具体原因待分析。但前面33个字节与通过 private key生成的public key相同。因此，直接截取前面33个字节
	b := base58.Decode(pubkeyStr)
	return b[:btcec.PubKeyBytesLenCompressed], len(b)
}

func ParseSignature(sigStr []byte) (*btcec.Signature, error) {
	return btcec.ParseSignature(sigStr, btcec.S256())
}

func GenerateKey() (*PrivateKey, error) {
	key, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return (*PrivateKey)((*btcec.PrivateKey)(key)), nil
}

func NewWIF(pk *PrivateKey) (*WIF, error) {
	w, err := btcutil.NewWIF((*btcec.PrivateKey)(pk), &chaincfg.MainNetParams, false)
	if err != nil {
		return nil, err
	}
	return (*WIF)(w), nil
}
