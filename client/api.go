package client

import (
	"errors"

	"github.com/weibocom/ipc/chain"
	"github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/store"
)

var (
	ErrAccountAlreadyExist = errors.New("account is already existed")
)

type Client interface {
	AccountCount() (uint32, error)
	CreateAccount(name string, meta string) (*model.Account, error)
	LookupAccount(name string) (*model.Account, error)
	GetAccounts(company string, offset int, limit int) ([]*model.Account, error)
	GetAccountPostCount(name string) (int, error)

	// chain
	Post(author string, mid int64, content []byte) (model.DNA, error)
	Verify(dna model.DNA) bool

	CheckSimilar(a, b model.DNA) (float64, error)
	LookupContent(dna model.DNA) (model.Content, error)
	LookupPost(author string, dna model.DNA) (*model.Post, error)
	LookupPostByMsgID(author string, mid int64) (*model.Post, error)
	LookupPostByDNA(dna model.DNA) (*model.Post, error)
	LookupPostByAuthor(author string, offset int, limit int) ([]*model.Post, error)
	GetLatestPost() (*model.Post, error)
	PostCount() (int, error)
	LookupSimilarPosts(dna string, keywords string, offset int, limit int) ([]*model.Post, error)

	Close() error
}

func NewClient(ipchain chain.Chain, store store.Store) (Client, error) {
	client := &client{
		ipchain: ipchain,
		store:   store,
		done:    make(chan struct{}),
	}

	return client, nil
}

type client struct {
	ipchain chain.Chain
	store   store.Store

	done chan struct{}
}

func (c *client) Close() error {
	close(c.done)
	_ = c.store.Close()
	return c.ipchain.Close()
}
