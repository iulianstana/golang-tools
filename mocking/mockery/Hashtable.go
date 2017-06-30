package mockery

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

// HashTable is the interface for a simple hash table
type HashTable interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}


type inMemoryHashTable struct {
	m map[string][]byte
}

func NewInMemoryHashTable() HashTable {
	return &inMemoryHashTable{m: make(map[string][]byte)}
}

func (i *inMemoryHashTable) Get(key string) ([]byte, error) {
	val, ok := i.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

func (i *inMemoryHashTable) Set(key string, val []byte) error {
	i.m[key] = val
	return nil
}
