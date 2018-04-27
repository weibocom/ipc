package signature

type Signature interface {
	Sign(privKeys [][]byte, digest []byte) ([][]byte, error)
	Verify(pubKeys [][]byte, digest []byte, sigs [][]byte) (bool, error)
}

func NewSignature() Signature {
	return &secp256k1{}
}
