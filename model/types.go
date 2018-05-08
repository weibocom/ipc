package model

import (
	"time"

	"github.com/weibocom/ipc/keys"
)

type Account struct {
	Name string
	WIF  *keys.WIF
}

type DNA []byte

func (dna DNA) ID() string {
	return string(dna)
}

type Member struct {
	Name       string
	ID         int64 // id in blockchain
	SigningKey string
	CreatedAt  time.Time
}

type Content []byte

type Post struct {
	Author  string
	Title   string
	Content []byte
	URI     string
	Digest  []byte
	DNA     DNA
}
