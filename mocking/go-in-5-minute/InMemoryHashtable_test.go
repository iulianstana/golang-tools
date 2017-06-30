package mocking

import "testing"

func TestBusinessLogic(t *testing.T) {
	// we create NewMemoryHashTable that match the hashtable interface
	ht := NewInMemoryHashTable()

	// insert new value in ht using businessLogic
	BusinessLogic(ht)

	// Test insertion
	val, err := ht.Get("hello")
	if err != nil {
		t.Fatalf("error on Get: %s", err)
	}
	if string(val) != "world" {
		t.Fatalf("expected 'world', got '%s'", string(val))
	}
}
