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

func (e cacheEntry[T]) IsExpired(tll time.Duration) bool {
	return time.Since(e.CreatedAt) > tll
}

type Cache[T any] struct {
	mu    sync.RWMutex
	store map[string]cacheEntry[T]
	ttl   time.Duration
}

func NewCache[T any](ttl time.Duration) *Cache[T] {
	if ttl == 0 {
		ttl = TTL
	}

	return &Cache[T]{
		mu:    sync.RWMutex{},
		store: make(map[string]cacheEntry[T]),
		ttl:   ttl,
	}
}

func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.store[key]
	if !exists || entry.IsExpired(c.ttl) {
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
