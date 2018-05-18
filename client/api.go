package client

import (
	"errors"

	"github.com/weibocom/ipc/interfaces"
	"github.com/weibocom/ipc/model"
	steemclient "github.com/weibocom/ipc/steem/client"
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

	Post(author string, mid int64, title string, content []byte, uri string, tags []string) (model.DNA, error)
	PostAsync(author string, mid int64, title string, content []byte, uri string, tags []string) (model.DNA, error)
	Verify(author string, dna model.DNA) (bool, error)
	CheckSimilar(a, b model.DNA) (float64, error)
	LookupContent(dna model.DNA) (model.Content, error)
	LookupPost(author string, dna model.DNA) (*model.Post, error)
	LookupPostByMsgID(author string, mid int64) (*model.Post, error)
	LookupPostByDNA(dna model.DNA) (*model.Post, error)
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
