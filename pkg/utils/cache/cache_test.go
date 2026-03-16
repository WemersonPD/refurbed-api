package utils_cache_test

import (
	utils_cache "assignment-backend/pkg/utils/cache"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var emptyTTL time.Duration

func Setup[T any](ttl time.Duration) *utils_cache.Cache[T] {
	return utils_cache.NewCache[T](ttl)
}

func TestCache_Get(t *testing.T) {
	t.Run("Successfully returning a cached response", func(t *testing.T) {
		cache := Setup[string](emptyTTL)
		cache.Set("foo", "bar")

		val, ok := cache.Get("foo")

		assert.True(t, ok, "Successfully check")
		assert.Equal(t, "bar", val, "Successfully returned")
	})

	t.Run("Missing key", func(t *testing.T) {
		cache := Setup[string](emptyTTL)
		cache.Set("foo", "bar")

		val, ok := cache.Get("bar")

		assert.False(t, ok, "Successfully check")
		assert.Empty(t, val, "No response")
	})

	t.Run("Expired key", func(t *testing.T) {
		shortTTL := 1 * time.Nanosecond
		cache := Setup[string](shortTTL)
		cache.Set("foo", "bar")
		time.Sleep(shortTTL)

		val, ok := cache.Get("foo")
		assert.False(t, ok, "Expired key")
		assert.Empty(t, val, "No response")
	})
}

func TestCache_Set(t *testing.T) {
	t.Run("Successfully setting a new entry", func(t *testing.T) {
		cache := Setup[string](emptyTTL)

		cache.Set("key", "value")

		val, ok := cache.Get("key")
		assert.True(t, ok)
		assert.Equal(t, "value", val)
	})

	t.Run("Overwriting an existing entry", func(t *testing.T) {
		cache := Setup[string](emptyTTL)

		cache.Set("key", "old")
		cache.Set("key", "new")

		val, ok := cache.Get("key")
		assert.True(t, ok)
		assert.Equal(t, "new", val)
	})

	t.Run("Concurrency check", func(t *testing.T) {
		cache := Setup[int](emptyTTL)

		cache.Set("a", 1)
		cache.Set("b", 2)
		cache.Set("c", 3)

		valA, okA := cache.Get("a")
		valB, okB := cache.Get("b")
		valC, okC := cache.Get("c")

		assert.True(t, okA)
		assert.Equal(t, 1, valA)
		assert.True(t, okB)
		assert.Equal(t, 2, valB)
		assert.True(t, okC)
		assert.Equal(t, 3, valC)
	})
}
