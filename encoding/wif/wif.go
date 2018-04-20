package wif

import "github.com/btcsuite/btcutil"

type WIF btcutil.WIF

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
