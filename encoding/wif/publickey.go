package wif

import (
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"github.com/weibocom/steem-rpc/config"
	"golang.org/x/crypto/ripemd160"
)

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
	prefix := config.GetAddressPrefix()
	key := str
	if strings.IndexAny(key, prefix) == 0 {
		key = key[len(prefix):]
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
		str = config.GetAddressPrefix() + str
	}

	return str
}
