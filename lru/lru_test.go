package lru

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if got := cache.Get(1); got != 1 {
		t.Errorf("got %d, want %d", got, 1)
	}
	cache.Put(3, 3)
	if got := cache.Get(3); got != 3 {
		t.Errorf("got %d, want %d", got, 3)
	}
	if got := cache.Get(2); got != -1 {
		t.Errorf("got %d, want %d", got, -1)
	}
	cache.Put(3, 5)
	if got := cache.Get(3); got != 5 {
		t.Errorf("got %d, want %d", got, 5)
	}
}
