package signature

// #cgo LDFLAGS: -lsecp256k1
// #include <stdlib.h>
// #include "signing.h"
import "C"
import (
	"bytes"
	"errors"
	"unsafe"
)

type secp256k1 struct {
}

func (s *secp256k1) Sign(privKeys [][]byte, digest []byte) ([][]byte, error) {
	// Sign.
	cDigest := C.CBytes(digest)
	defer C.free(cDigest)

	cKeys := make([]unsafe.Pointer, 0, len(privKeys))
	for _, key := range privKeys {
		cKeys = append(cKeys, C.CBytes(key))
	}
	defer func() {
		for _, cKey := range cKeys {
			C.free(cKey)
		}
	}()

	sigs := make([][]byte, 0, len(privKeys))
	for _, cKey := range cKeys {
		var (
			signature [64]byte
			recid     C.int
		)

		code := C.sign_transaction(
			(*C.uchar)(cDigest), (*C.uchar)(cKey), (*C.uchar)(&signature[0]), &recid)
		if code == 0 {
			return nil, errors.New("sign_transaction returned 0")
		}

		sig := make([]byte, 65)
		sig[0] = byte(recid)
		copy(sig[1:], signature[:])
		sigs = append(sigs, sig)
	}
	return sigs, nil
}

func (s *secp256k1) Verify(pubKeys [][]byte, digest []byte, sigs [][]byte) (bool, error) {
	cDigest := C.CBytes(digest)
	defer C.free(cDigest)

	// Make sure to free memory.
	cSigs := make([]unsafe.Pointer, 0, len(sigs))
	defer func() {
		for _, cSig := range cSigs {
			C.free(cSig)
		}
	}()

	// Collect verified public keys.
	pubKeysFound := make([][]byte, len(pubKeys))
	for i, sig := range sigs {
		recoverParameter := sig[0] - 27 - 4
		sig = sig[1:]

		cSig := C.CBytes(sig)
		cSigs = append(cSigs, cSig)

		var publicKey [33]byte

		code := C.verify_recoverable_signature(
			(*C.uchar)(cDigest),
			(*C.uchar)(cSig),
			(C.int)(recoverParameter),
			(*C.uchar)(&publicKey[0]),
		)
		if code == 1 {
			pubKeysFound[i] = publicKey[:]
		}
	}

	for i := range pubKeys {
		if !bytes.Equal(pubKeysFound[i], pubKeys[i]) {
			return false, nil
		}
	}
	return true, nil
}
