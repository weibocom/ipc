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

// String 1. serialize bytes; 2. 末尾增加4个字节的ripemd160的hash值 3. 用base56编码
func (p *PublicKey) String() string {
	ser := p.Serialize()
	hash := ripemd160.New()
	hash.Write(ser)
	sum := hash.Sum(nil)[:4]

	serWithSum := append(ser, sum...)

	return base58.Encode(serWithSum)
}

// ParsePublicKey returns the public key associated with the given public-base58-formatted stringent key
// in the 33-byte compressed format.
// pubkeyStr  不包含任何前缀。比如STM等。
func ParsePubKeyBase58(pubkeyStr string) ([]byte, int) {
	// TODO: 该函数返回的字节数并不是33，而是37. 具体原因待分析。但前面33个字节与通过 private key生成的public key相同。因此，直接截取前面33个字节
	b := base58.Decode(pubkeyStr)
	return b[:btcec.PubKeyBytesLenCompressed], len(b)
}
