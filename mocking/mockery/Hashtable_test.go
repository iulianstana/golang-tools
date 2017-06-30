package mockery

import "testing"
import "mocking/mockery/mocks"
import . "github.com/stretchr/testify/mock"

func TestBusinessLogic(t *testing.T) {
	// we create NewMemoryHashTable that match the hashtable interface
	ht := NewInMemoryHashTable()

	// insert new value in ht using businessLogic
	ht.Set("hello", []byte("world"))

	// Test insertion
	val, err := ht.Get("hello")
	if err != nil {
		t.Fatalf("error on Get: %s", err)
	}
	if string(val) != "world" {
		t.Fatalf("expected 'world', got '%s'", string(val))
	}
}

func TestBusinessLogicMockery(t *testing.T) {
	// we create NewMemoryHashTable that match the hashtable interface
	ht := mocks.HashTable{}

	// insert new value in ht using businessLogic
	ht.On("Get", AnythingOfType("string")).Return(func(s string) []byte {
	    return []byte("Mock string")
	}, nil)

	// Test insertion
	val, err := ht.Get("hello")
	if err != nil {
		t.Fatalf("error on Get: %s", err)
	}
	if string(val) != "Mock string" {
		t.Fatalf("expected 'Mock string', got '%s'", string(val))
	}
}
