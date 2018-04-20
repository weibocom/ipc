package steem

import "github.com/weibocom/steem-rpc/config"

type Chain struct {
	ID string
}

var SteemChain = &Chain{
	ID: config.GetChainID(),
}
