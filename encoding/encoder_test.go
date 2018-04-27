package encoding

import (
	// Stdlib
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestDecode(t *testing.T) {
	expectedHex := "9115fbb617ba0385a5bffba37230f02a9bc242e940e9f51a30bb2a451fb041e0"

	var b bytes.Buffer
	encoder := NewEncoder(&b)

	var refBlockNumber uint16 = 12699

	if err := encoder.EncodeNumber(refBlockNumber); err != nil {
		t.Error(err)
	}

	hash := sha256.Sum256(b.Bytes())

	serializedHex := hex.EncodeToString(hash[:])

	if serializedHex != expectedHex {
		t.Errorf("expected %v, got %v", expectedHex, serializedHex)
	}
}
