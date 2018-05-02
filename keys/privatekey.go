package keys

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
)

type PrivateKey btcec.PrivateKey

func (p *PrivateKey) pk() *btcec.PrivateKey {
	return (*btcec.PrivateKey)(p)
}

func (p *PrivateKey) Public() *PublicKey {
	return (*PublicKey)(p.pk().PubKey())
}

// Serialize can be used to turn private key into a raw private key (32 bytes).
func (p *PrivateKey) Serialize() []byte {
	return p.pk().Serialize()
}

func (p *PrivateKey) FromBytes(b []byte) {
	pk, _ := btcec.PrivKeyFromBytes(btcec.S256(), b)
	(*p) = (PrivateKey)(*pk)
}

func (p *PrivateKey) HexString() string {
	return hex.EncodeToString(p.Serialize())
}

func GenerateKey() (*PrivateKey, error) {
	key, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return (*PrivateKey)((*btcec.PrivateKey)(key)), nil
}
