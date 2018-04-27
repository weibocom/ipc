package wif

import "github.com/btcsuite/btcd/btcec"

type PrivateKey btcec.PrivateKey

func (p *PrivateKey) pk() *btcec.PrivateKey {
	return (*btcec.PrivateKey)(p)
}

// Serialize can be used to turn private key into a raw private key (32 bytes).
func (p *PrivateKey) Serialize() []byte {
	return p.pk().Serialize()
}
