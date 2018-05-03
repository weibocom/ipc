package store

import "sync"

type Store interface {
	Save(storeType string, key string, value []byte) error
	Load(storeType string, key string) ([]byte, error)
	Exist(storeType string, key string) (bool, error)
	Close() error
}

func generateKey(prefix string, storeType string, key string) string {
	return prefix + "-" + storeType + "-" + key
}

var _ Store = &MemStore{}

// MemStore implements the Store in memory for testing purposes.
type MemStore struct {
	prefix string
	data   map[string][]byte
	mu     sync.RWMutex
}

func NewMemStore(prefix string) *MemStore {
	return &MemStore{
		prefix: prefix,
		data:   make(map[string][]byte),
	}
}

func (s *MemStore) Save(storeType string, key string, value []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[generateKey(s.prefix, storeType, key)] = value
	return nil
}

func (s *MemStore) Load(storeType string, key string) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[generateKey(s.prefix, storeType, key)], nil
}

func (s *MemStore) Exist(storeType string, key string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exist := s.data[generateKey(s.prefix, storeType, key)]
	return exist, nil
}

func (s *MemStore) Close() error {
	return nil
}
