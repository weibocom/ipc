package store

import (
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/weibocom/ipc/model"
)

var _ Store = &MemStore{}

// MemStore implements the Store in memory for testing purposes.
type MemStore struct {
	prefix    string
	accounts  map[string]*model.Account
	members   map[string]*model.Member
	posts     map[string]*model.Post
	posts2    map[string]*model.Post
	lastPost  *model.Post
	mu        sync.RWMutex
	postCount uint64
}

func NewMemStore(prefix string) *MemStore {
	return &MemStore{
		prefix:   prefix,
		accounts: make(map[string]*model.Account),
		members:  make(map[string]*model.Member),
		posts:    make(map[string]*model.Post),
		posts2:   make(map[string]*model.Post),
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

func (s *MemStore) GetAccounts(company string, offset int, limit int) ([]*model.Account, error) {
	return nil, ErrNotImplemented
}

func (s *MemStore) GetAccountCount() (int, error) {
	return 0, ErrNotImplemented
}

func (s *MemStore) GetPostCount() (int, error) {
	count := atomic.LoadUint64(&s.postCount)
	return int(count), ErrNotImplemented
}

func (s *MemStore) GetAccountPostCount(name string) (int, error) {
	return 0, ErrNotImplemented
}

func (s *MemStore) SavePost(p *model.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.posts[p.DNA] = p
	s.posts2[p.Author+"-"+strconv.FormatInt(p.MSGID, 10)] = p
	s.lastPost = p

	atomic.AddUint64(&s.postCount, 1)
	return nil
}

func (s *MemStore) LoadPost(dna model.DNA) (*model.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.posts[dna.String()]
	if ok {
		return v, nil
	}
	return nil, ErrNonExist
}

func (s *MemStore) ExistPost(dna model.DNA) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exist := s.posts[dna.String()]
	return exist, nil
}

func (s *MemStore) GetLatestPost() (*model.Post, error) {
	return s.lastPost, nil
}

func (s *MemStore) GetPostByMsgID(author string, mid int64) (*model.Post, error) {
	return s.posts2[author+"-"+strconv.FormatInt(mid, 10)], nil
}

func (s *MemStore) GetPostByDNA(dna model.DNA) (*model.Post, error) {
	return s.posts[dna.String()], nil
}

func (s *MemStore) GetPostByAuthor(author string, offset int, limit int) ([]*model.Post, error) {
	return nil, ErrNotImplemented
}

func (s *MemStore) LookupSimilarPosts(dna string, keywords string, offset int, limit int) ([]*model.Post, error) {
	return nil, ErrNotImplemented
}

func (s *MemStore) Close() error {
	return nil
}
