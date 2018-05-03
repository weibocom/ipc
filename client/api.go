package client

import (
	"github.com/weibocom/ipc/interfaces"
	steemclient "github.com/weibocom/ipc/steem/client"
	"github.com/weibocom/ipc/store"
)

const (
	AccountStoreType = "account"
	PostStoreType    = "post"
	MemberStoreType  = "member"
)

type Client interface {
	CreateAccount(name string, meta string) (*Account, error)
	Post(author string, title string, content []byte, uri string, tags []string) (DNA, error)
	LookupContent(dna DNA) (Content, error)
	Verify(author string, dna DNA) (bool, error)
	// CheckSimilar(a, b DNA) (int, error)
	// Members() (error, []Member)
	// AddMember(m Member) error
	// RemoveMember(m Member) error
	Close() error
}

func NewClient(cc interfaces.CallCloser, store store.Store) (Client, error) {
	steem, err := steemclient.NewClient(cc)
	if err != nil {
		return nil, err
	}
	return &client{steem: steem, store: store}, nil
}

type client struct {
	steem *steemclient.Client
	store store.Store
}

func (c *client) Close() error {
	return c.steem.Close()
}
