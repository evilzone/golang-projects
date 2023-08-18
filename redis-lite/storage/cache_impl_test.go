package storage

import (
	"errors"
	"testing"
)

func TestNewInMemoryStorage(t *testing.T) {
	cache := NewInMemoryStorage()

	t.Run("get key after set returns value ", func(t *testing.T) {
		cache.Set("key", []byte("value"))
		value, err := cache.Get("key")

		if err != nil {
			t.Errorf("Expected error to be nil %v", err)
		}

		if string(value) != "value" {
			t.Errorf("Expected value to be 'value' but got %v", value)
		}
	})

	t.Run("get key doesn't exist returns error ", func(t *testing.T) {
		_, err := cache.Get("key1")

		if !errors.Is(err, ErrKeyNotFound) {
			t.Errorf("Expected error to be ErrKeyNotFound but found %v", err)
		}
	})

	t.Run("delete key removes key ", func(t *testing.T) {
		cache.Set("test", []byte("val"))
		cache.Delete("test")
		_, err := cache.Get("test")

		if !errors.Is(err, ErrKeyNotFound) {
			t.Errorf("Expected error to be ErrKeyNotFound but found %v", err)
		}
	})
}
