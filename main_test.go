package cache

import "testing"

func TestCache2(t *testing.T) {
	cache := NewCache()
	cache.Put("a", "a1")
	cache.Put("b", "b1")
	cache.Put("c", "c1")
	cache.Put("d", "d1")
	v := cache.Keys()
}
