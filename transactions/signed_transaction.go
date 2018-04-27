// +build !nosigning

package transactions

import (
	// Stdlib
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"time"

	// RPC

	"github.com/weibocom/steem-rpc/encoding/transaction"
	"github.com/weibocom/steem-rpc/signature"
	"github.com/weibocom/steem-rpc/steem"
	"github.com/weibocom/steem-rpc/steem/types"

	// Vendor
	"github.com/pkg/errors"
)

// // #cgo LDFLAGS: -lsecp256k1
// // #include <stdlib.h>
// // #include "signing.h"
// import "C"

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

func (tx *SignedTransaction) Serialize() ([]byte, error) {
	var b bytes.Buffer
	encoder := transaction.NewEncoder(&b)

	if err := encoder.Encode(tx.Transaction); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (tx *SignedTransaction) Digest(chain *steem.Chain) ([]byte, error) {
	var msgBuffer bytes.Buffer

	// Write the chain ID.
	rawChainID, err := chain.DecodeID()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode chain ID: %v", chain.ID)
	}

	if _, err := msgBuffer.Write(rawChainID); err != nil {
		return nil, errors.Wrap(err, "failed to write chain ID")
	}

	// Write the serialized transaction.
	rawTx, err := tx.Serialize()
	if err != nil {
		return nil, err
	}

	if _, err := msgBuffer.Write(rawTx); err != nil {
		return nil, errors.Wrap(err, "failed to write serialized transaction")
	}

	// Compute the digest.
	digest := sha256.Sum256(msgBuffer.Bytes())
	return digest[:], nil
}

// 基于C实现的签名，与CVerify对应
func (tx *SignedTransaction) Sign(privKeys [][]byte, chain *steem.Chain) error {
	digest, err := tx.Digest(chain)
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

func (tx *SignedTransaction) Verify(pubKeys [][]byte, chain *steem.Chain) (bool, error) {
	// Compute the digest, again.
	digest, err := tx.Digest(chain)
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
