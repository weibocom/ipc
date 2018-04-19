package wif

import (
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/weibocom/steem-rpc/steem"
	"golang.org/x/crypto/ripemd160"
)

type WIF btcutil.WIF

type PrivateKey btcec.PrivateKey

type PublicKey btcec.PublicKey

func (p *PublicKey) pk() *btcec.PublicKey {
	return (*btcec.PublicKey)(p)
}

// GetPublicKey returns the public key associated with the given WIF
// in the 33-byte compressed format.
func (p *PublicKey) Serialize() []byte {
	return p.pk().SerializeCompressed()
}

// 输入的是BASE58编码的字符串
func (p *PublicKey) From(str string) error {
	key := str
	if strings.IndexAny(key, steem.STEEM_ADDRESS_PREFIX) == 0 {
		key = key[len(steem.STEEM_ADDRESS_PREFIX):]
	}
	b := base58.Decode(key)
	return p.FromBytes(b)
}

func (p *PublicKey) FromBytes(b []byte) error {
	// 最后4个字节为ripe160校验。直接忽略
	trunc := b[:btcec.PubKeyBytesLenCompressed]
	key, err := btcec.ParsePubKey(trunc, btcec.S256())
	if err != nil {
		return err
	}
	*p = (PublicKey)(*key)
	return nil
}

func (p *PublicKey) String(prefix bool) string {
	ser := p.Serialize()
	hash := ripemd160.New()
	hash.Write(ser)
	sum := hash.Sum(nil)[:4]

	serWithSum := append(ser, sum...)

	str := base58.Encode(serWithSum)
	if prefix {
		str = steem.STEEM_ADDRESS_PREFIX + str
	}

	return str
}

func (w *WIF) PrivateKey() *PrivateKey {
	return (*PrivateKey)(w.PrivKey)
}

func (w *WIF) PublicKey() *PublicKey {
	return (*PublicKey)(w.PrivKey.PubKey())
}

// Serialize can be used to turn WIF into a raw private key (32 bytes).
func (w *WIF) Serialize() []byte {
	return w.PrivKey.Serialize()
}

func (w *WIF) String() string {
	return ((*btcutil.WIF)(w)).String()
}
