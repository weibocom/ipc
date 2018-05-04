package store

import (
	"github.com/bradfitz/gomemcache/memcache"
)

// MemcacheStore implements the Store with memcached.
type MemcacheStore struct {
	prefix string
	mc     *memcache.Client
}

var _ Store = &MemcacheStore{}

func NewMemcacheStore(prefix string, server ...string) *MemcacheStore {
	mc := memcache.New(server...)
	return &MemcacheStore{prefix: prefix, mc: mc}
}

func (s *MemcacheStore) Client() *memcache.Client {
	return s.mc
}

func (s *MemcacheStore) Save(storeType string, key string, value []byte) error {
	key = generateKey(s.prefix, storeType, key)
	return s.mc.Set(&memcache.Item{
		Key:   key,
		Value: value,
	})
}

func (s *MemcacheStore) Load(storeType string, key string) ([]byte, error) {
	key = generateKey(s.prefix, storeType, key)
	item, err := s.mc.Get(key)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, ErrNonExist
		}
		return nil, err
	}

	return item.Value, nil
}

func (s *MemcacheStore) Exist(storeType string, key string) (bool, error) {
	key = generateKey(s.prefix, storeType, key)
	_, err := s.mc.Get(key)
	if err == memcache.ErrCacheMiss {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *MemcacheStore) Close() error {
	return nil
}
