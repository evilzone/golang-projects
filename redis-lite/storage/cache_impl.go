package storage

import (
	"errors"
	"sync"
)

type InMemoryStorage struct {
	mu   sync.Mutex
	data map[string][]byte
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data: make(map[string][]byte),
	}
}

func (i *InMemoryStorage) Set(key string, value []byte) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.data[key] = value
}

func (i *InMemoryStorage) Get(key string) ([]byte, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	val, ok := i.data[key]

	if !ok {
		return nil, errors.New("key not found")
	}
	return val, nil
}

func (i *InMemoryStorage) Delete(key string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	delete(i.data, key)
}
