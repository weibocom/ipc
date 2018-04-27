package types

import (
	// RPC
	"github.com/weibocom/steem-rpc/encoding"

	// Vendor
	"github.com/pkg/errors"
)

// Transaction represents a blockchain transaction.
type Transaction struct {
	RefBlockNum    UInt16            `json:"ref_block_num"`
	RefBlockPrefix UInt32            `json:"ref_block_prefix"`
	Expiration     *TimePointSeconds `json:"expiration"`
	Operations     Operations        `json:"operations"`
	Signatures     []string          `json:"signatures"`
}

// MarshalTransaction implements transaction.Marshaller interface.
func (tx Transaction) Marshal(encoder *encoding.Encoder) error {
	if len(tx.Operations) == 0 {
		return errors.New("no operation specified")
	}

	enc := encoding.NewRollingEncoder(encoder)

	enc.Encode(tx.RefBlockNum)
	enc.Encode(tx.RefBlockPrefix)
	enc.Encode(tx.Expiration)

	enc.EncodeUVarint(uint64(len(tx.Operations)))
	for _, op := range tx.Operations {
		enc.Encode(op)
	}

	// Extensions are not supported yet.
	enc.EncodeUVarint(0)

	return enc.Err()
}

// PushOperation can be used to add an operation into the transaction.
func (tx *Transaction) PushOperation(op Operation) {
	tx.Operations = append(tx.Operations, op)
}
