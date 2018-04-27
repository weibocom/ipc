package steem

import (
	"encoding/hex"

	"github.com/weibocom/ipc/config"
)

type Chain struct {
	ID string
}

var SteemChain = &Chain{
	ID: config.GetChainID(),
}

// Decode decode to hex format
func (c Chain) DecodeID() ([]byte, error) {
	return hex.DecodeString(c.ID)
}
