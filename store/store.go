package store

import (
	"errors"

	"github.com/weibocom/ipc/model"
)

var (
	ErrNonExist       = errors.New("not exist")
	ErrExist          = errors.New("already exist")
	ErrNotImplemented = errors.New("not implemented yet")
)

type Store interface {
	Account
	Member
	Post
	Close() error
}

type Account interface {
	ExistAccount(name string) (bool, error)
	SaveAccount(a *model.Account) error
	LoadAccount(name string) (*model.Account, error)
}

type Member interface {
	ExistMember(name string) (bool, error)
	SaveMember(m *model.Member) error
	LoadMember(name string) (*model.Member, error)
}

type Post interface {
	ExistPost(dna model.DNA) (bool, error)
	SavePost(p *model.Post) error
	LoadPost(dna model.DNA) (*model.Post, error)
	GetPosts(author string, afterDNA model.DNA, limit int) ([]*model.Post, error)
	GetLatestPost() (*model.Post, error)
	GetPostByURI(author string, uri string) (*model.Post, error)
}
