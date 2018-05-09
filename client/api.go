package client

import (
	"github.com/weibocom/ipc/interfaces"
	"github.com/weibocom/ipc/model"
	steemclient "github.com/weibocom/ipc/steem/client"
	"github.com/weibocom/ipc/store"
)

type Client interface {
	AccountCount() (uint32, error)
	CreateAccount(name string, meta string) (*model.Account, error)

	Post(author string, title string, content []byte, uri string, tags []string) (model.DNA, error)
	Verify(author string, dna model.DNA) (bool, error)
	CheckSimilar(a, b model.DNA) (float64, error)
	LookupContent(dna model.DNA) (model.Content, error)
	LookupPost(author string, dna model.DNA) (*model.Post, error)
	LookupPostByURI(author string, uri string) (*model.Post, error)
	GetPosts(author string, afterDNA model.DNA, limit int) ([]*model.Post, error)
	GetLatestPost() (*model.Post, error)

	Members() ([]*model.Member, error)
	AddMember(name string) (*model.Member, error)
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
	_ = c.store.Close()
	return c.steem.Close()
}
