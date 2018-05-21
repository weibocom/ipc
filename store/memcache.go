package store

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/util"
)

// MemcacheStore implements the Store with memcached.
type MemcacheStore struct {
	prefix string
	mc     *memcache.Client
}

var _ Store = &MemcacheStore{}

func generateKey(prefix string, storeType string, key string) string {
	return prefix + "-" + storeType + "-" + key
}

func NewMemcacheStore(prefix string, server ...string) *MemcacheStore {
	mc := memcache.New(server...)
	return &MemcacheStore{prefix: prefix, mc: mc}
}

func (s *MemcacheStore) Client() *memcache.Client {
	return s.mc
}

func (s *MemcacheStore) SaveAccount(a *model.Account) error {
	v, err := util.ToJSON(a)
	if err != nil {
		return err
	}
	key := generateKey(s.prefix, "account", a.Name)
	return s.mc.Set(&memcache.Item{
		Key:   key,
		Value: v,
	})

	return nil
}

func (s *MemcacheStore) LoadAccount(name string) (*model.Account, error) {
	key := generateKey(s.prefix, "account", name)
	item, err := s.mc.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, ErrNonExist
		}
		return nil, err
	}

	a := &model.Account{}
	err = util.FromJSON(item.Value, a)
	return a, err
}

func (s *MemcacheStore) ExistAccount(name string) (bool, error) {
	a, err := s.LoadAccount(name)
	if err == ErrNonExist {
		return false, nil
	}

	return a != nil, err
}

func (s *MemcacheStore) GetAccounts(company string, offset int, limit int) ([]*model.Account, error) {
	return nil, ErrNotImplemented
}

func (s *MemcacheStore) SaveMember(m *model.Member) error {
	v, err := util.ToJSON(m)
	if err != nil {
		return err
	}
	key := generateKey(s.prefix, "member", m.Name)
	return s.mc.Set(&memcache.Item{
		Key:   key,
		Value: v,
	})

	return nil
}

func (s *MemcacheStore) LoadMember(name string) (*model.Member, error) {
	key := generateKey(s.prefix, "member", name)
	item, err := s.mc.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, ErrNonExist
		}
		return nil, err
	}

	a := &model.Member{}
	err = util.FromJSON(item.Value, a)
	return a, err
}

func (s *MemcacheStore) ExistMember(name string) (bool, error) {
	m, err := s.LoadMember(name)
	if err == ErrNonExist {
		return false, nil
	}
	return m != nil, err
}

func (s *MemcacheStore) GetPostCount() (int, error) {
	return 0, ErrNotImplemented
}

func (s *MemcacheStore) SavePost(p *model.Post) error {
	v, err := util.ToJSON(p)
	if err != nil {
		return err
	}
	key := generateKey(s.prefix, "post", p.DNA)
	return s.mc.Set(&memcache.Item{
		Key:   key,
		Value: v,
	})

	return nil
}

func (s *MemcacheStore) LoadPost(dna model.DNA) (*model.Post, error) {
	key := generateKey(s.prefix, "post", dna.ID())
	item, err := s.mc.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, ErrNonExist
		}
		return nil, err
	}

	a := &model.Post{}
	err = util.FromJSON(item.Value, a)
	return a, err
}

func (s *MemcacheStore) ExistPost(dna model.DNA) (bool, error) {
	p, err := s.LoadPost(dna)
	if err == ErrNonExist {
		return false, nil
	}
	return p != nil, err
}

func (s *MemcacheStore) GetLatestPost() (*model.Post, error) {
	return nil, ErrNotImplemented
}

func (s *MemcacheStore) GetPostByMsgID(author string, mid int64) (*model.Post, error) {
	return nil, ErrNotImplemented
}

func (s *MemcacheStore) GetPostByDNA(dna model.DNA) (*model.Post, error) {
	return nil, ErrNotImplemented
}

func (s *MemcacheStore) GetPostByAuther(author string, offset int, limit int) ([]*model.Post, error) {
	return nil, ErrNotImplemented
}

func (s *MemcacheStore) Close() error {
	return nil
}
