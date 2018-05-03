package client

import "github.com/weibocom/ipc/keys"

type Account struct {
	Name string
	WIF  *keys.WIF
}

type DNA []byte

func (dna DNA) ID() string {
	return string(dna)
}

type Member struct{}

type Content struct{}
