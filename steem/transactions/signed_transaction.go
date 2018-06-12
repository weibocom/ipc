package transactions

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"time"

	// RPC

	"github.com/weibocom/ipc/encoding"
	"github.com/weibocom/ipc/signature"
	"github.com/weibocom/ipc/steem/types"
	// Vendor
)

type SignedTransaction struct {
	*types.Transaction
}

func NewSignedTransaction(tx *types.Transaction) *SignedTransaction {
	if tx.Expiration == nil {
		expiration := time.Now().Add(30 * time.Second)
		tx.Expiration = types.NewTimePointSeconds(expiration)
	}

	return &SignedTransaction{tx}
}

func (tx *SignedTransaction) Digest(chainID string) ([]byte, error) {
	var b bytes.Buffer
	encoder := encoding.NewRollingEncoder(encoding.NewEncoder(&b))
	encoder.Encode(chainID)
	encoder.Encode(tx.Transaction)

	if encoder.Err() != nil {
		return nil, encoder.Err()
	}

	// Compute the digest.
	digest := sha256.Sum256(b.Bytes())
	return digest[:], nil
}

// 基于C实现的签名，与CVerify对应
func (tx *SignedTransaction) Sign(privKeys [][]byte, chainID string) error {
	digest, err := tx.Digest(chainID)
	if err != nil {
		return err
	}

	s := signature.NewSignature()
	sigs, err := s.Sign(privKeys, digest)
	if err != nil {
		return err
	}

	// Set the signature array in the transaction.
	sigsHex := make([]string, 0, len(sigs))
	for _, sig := range sigs {
		sigsHex = append(sigsHex, hex.EncodeToString(sig))
	}

	tx.Transaction.Signatures = sigsHex
	return nil
}

func (tx *SignedTransaction) Verify(pubKeys [][]byte, chainID string) (bool, error) {
	// Compute the digest, again.
	digest, err := tx.Digest(chainID)
	if err != nil {
		return false, err
	}

	sigs := make([][]byte, 0, len(tx.Signatures))
	for _, sigStr := range tx.Signatures {
		sig, err := hex.DecodeString(sigStr)
		if err != nil {
			return false, err
		}
		sigs = append(sigs, sig)
	}

	s := signature.NewSignature()
	return s.Verify(pubKeys, digest, sigs)
}
