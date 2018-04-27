package chain

import (
	"encoding/hex"

	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/encoding"
)

type ChainID string

type Chain struct {
	id ChainID
}

var MainChain = &Chain{
	id: ChainID(config.GetChainID()),
}

// Decode decode to hex format
func (c Chain) ID() ChainID {
	return c.id
}

func (c ChainID) String() string {
	return string(c)
}

func (c *ChainID) Serialize() ([]byte, error) {
	return hex.DecodeString(string(c.String()))
}

func (c ChainID) Marshal(encoder *encoding.Encoder) error {
	b, err := c.Serialize()
	if err == nil {
		err = encoder.Encode(b)
	}
	return err
}
