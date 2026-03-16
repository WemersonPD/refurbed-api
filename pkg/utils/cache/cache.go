package utils_cache

import (
	"sync"
	"time"
)

type cacheEntry[T any] struct {
	Data      T
	CreatedAt time.Time
}

const TTL = 30 * time.Second

func (e cacheEntry[T]) IsExpired() bool {
	return time.Since(e.CreatedAt) > TTL
}

type Cache[T any] struct {
	mu    sync.RWMutex
	store map[string]cacheEntry[T]
}

func NewCache[T any]() *Cache[T] {
	return &Cache[T]{
		store: make(map[string]cacheEntry[T]),
	}
}

func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.store[key]
	if !exists || entry.IsExpired() {
		var zero T
		return zero, false
	}

	return entry.Data, true
}

func (c *Cache[T]) Set(key string, data T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = cacheEntry[T]{
		Data:      data,
		CreatedAt: time.Now(),
	}
}
