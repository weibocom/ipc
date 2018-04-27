package types

import (
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/encoding"
)

// PublicKey base58编码后的字符串
type PublicKey string

func (p PublicKey) String() string {
	return string(p)
}

func (p PublicKey) Bytes() []byte {
	prefix := config.GetAddressPrefix()
	key := p.String()
	if strings.IndexAny(key, prefix) == 0 {
		key = key[len(prefix):]
	}
	b := base58.Decode(key)
	// TODO 没有校验后面4个字节的hash.
	return b[:btcec.PubKeyBytesLenCompressed]
}

func (p PublicKey) Marshal(encoder *encoding.Encoder) error {
	return encoder.Encode(p.Bytes())
}
