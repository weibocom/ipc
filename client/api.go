package client

import (
	"github.com/weibocom/ipc/interfaces"
	steemclient "github.com/weibocom/ipc/steem/client"
)

type Client interface {
	// CreateAccount(name string, meta interface{}) (*Account, error)
	CreateAccount(name string, meta string) (*Account, error)
	// Post(author string, title string, content []byte, abstracts []string, url string) (*DNA, error)
	// LookupContent(dna DNA) (*Content, error)
	// Verify(author string, dna DNA) (error, bool)
	// CheckSimilar(a, b DNA) (int, error)
	// Members() (error, []Member)
	// AddMember(m Member) error
	// RemoveMember(m Member) error
	Close() error
}

func NewClient(cc interfaces.CallCloser) (Client, error) {
	steem, err := steemclient.NewClient(cc)
	if err != nil {
		return nil, err
	}
	return &client{steem}, nil
}

type client struct {
	steem *steemclient.Client
}

func (c *client) Close() error {
	return c.steem.Close()
}
