package store

import (
	"sync"

	"github.com/weibocom/ipc/model"
)

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
	s.posts[p.DNA] = p
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

func (s *MemStore) GetPosts(author string, afterDNA model.DNA, limit int) ([]*model.Post, error) {
	return nil, ErrNotImplemented
}

func (s *MemStore) GetLatestPost() (*model.Post, error) {
	return nil, ErrNotImplemented
}

func (s *MemStore) GetPostByURI(author string, uri string) (*model.Post, error) {
	return nil, ErrNotImplemented
}

func (s *MemStore) Close() error {
	return nil
}
