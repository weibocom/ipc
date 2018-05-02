package signature

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weibocom/ipc/keys"
)

func TestSecp256k1(t *testing.T) {
	s := &secp256k1{}
	digest := []byte{58, 254, 69, 113, 4, 51, 129, 208, 5, 4, 212, 188, 77, 2, 237, 129, 249, 176, 113, 67, 23, 237, 201, 205, 28, 5, 7, 11, 29, 236, 223, 174}
	prvKey := []byte{155, 158, 2, 96, 141, 53, 4, 127, 4, 146, 187, 173, 184, 222, 184, 90, 94, 11, 143, 220, 102, 29, 99, 236, 240, 62, 234, 198, 140, 225, 195, 254}
	pubKey := []byte{2, 245, 118, 6, 5, 213, 177, 221, 175, 173, 204, 248, 217, 76, 77, 47, 237, 143, 103, 223, 133, 68, 31, 8, 144, 67, 34, 166, 86, 126, 192, 196, 34}

	sigs, err := s.Sign([][]byte{prvKey}, digest)
	require.NoError(t, err, "sign digest")

	pass, err := s.Verify([][]byte{pubKey}, digest, sigs)
	require.NoError(t, err, "verify signature")

	assert.True(t, pass, "verify signature")
}

func TestSecp256k1WIF(t *testing.T) {
	// test by witness
	s := &secp256k1{}
	wifStr := "5JWHY5DxTF6qN5grTtChDCYBmWHfY9zaSsw4CxEKN5eZpH9iBma"
	w, err := wif.DecodeWIF(wifStr)
	require.NoError(t, err, "decode wif:%s", wifStr)
	privKey := w.PrivateKey().Serialize()
	pubKey := w.PublicKey().Serialize()
	digestArray := sha256.Sum256([]byte(wifStr))
	digest := digestArray[:]
	sigs, err := s.Sign([][]byte{privKey}, digest)
	require.NoError(t, err, "sign digest")
	require.Len(t, sigs, 1, "only by one private key signed")
	assert.Len(t, sigs[0], 65, "signature length")
	pass, err := s.Verify([][]byte{pubKey}, digest, sigs)
	require.NoError(t, err, "verify digest")
	assert.True(t, pass, "verify signature")
}
