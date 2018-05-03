package client

import "github.com/weibocom/ipc/keys"

type Account struct {
	Name string
	WIF  *keys.WIF
}

type DNA struct{}

func (dna *DNA) ID() string {
	return ""
}

type Member struct{}

type Content struct{}
