package store

import (
	"errors"
	"sync"

	"github.com/weibocom/ipc/model"
)

var (
	ErrNonExist = errors.New("not exist")
	ErrExist    = errors.New("already exist")
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
}

var _ Store = &MemStore{}

// MemStore implements the Store in memory for testing purposes.
type MemStore struct {
	prefix   string
	accounts map[string]*model.Account
	members  map[string]*model.Member
	posts    map[string]*model.Post
	mu       sync.RWMutex
}

func NewMemStore(prefix string) *MemStore {
	return &MemStore{
		prefix:   prefix,
		accounts: make(map[string]*model.Account),
		members:  make(map[string]*model.Member),
		posts:    make(map[string]*model.Post),
	}
}

func (s *MemStore) SaveAccount(a *model.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.accounts[a.Name] = a
	return nil
}

func (s *MemStore) LoadAccount(name string) (*model.Account, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.accounts[name]
	if ok {
		return v, nil
	}
	return nil, ErrNonExist
}

func (s *MemStore) ExistAccount(name string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exist := s.accounts[name]
	return exist, nil
}

func (s *MemStore) SaveMember(m *model.Member) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.members[m.Name] = m
	return nil
}

func (s *MemStore) LoadMember(name string) (*model.Member, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.members[name]
	if ok {
		return v, nil
	}
	return nil, ErrNonExist
}

func (s *MemStore) ExistMember(name string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exist := s.members[name]
	return exist, nil
}

func (s *MemStore) SavePost(p *model.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.posts[p.DNA.ID()] = p
	return nil
}

func (s *MemStore) LoadPost(dna model.DNA) (*model.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.posts[dna.ID()]
	if ok {
		return v, nil
	}
	return nil, ErrNonExist
}

func (s *MemStore) ExistPost(dna model.DNA) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exist := s.posts[dna.ID()]
	return exist, nil
}
func (s *MemStore) Close() error {
	return nil
}
