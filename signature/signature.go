package signature

type Signature interface {
	Sign(pk []byte, digest []byte) ([]byte, error)
	Verify(pubKey []byte, digest []byte, sig []byte) (bool, error)
}
