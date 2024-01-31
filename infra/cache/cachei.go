package cache

import "context"

type BaseCache interface {
	Exist(ctx context.Context, key []byte) (bool, error)
}

type KVCacheI interface {
	BaseCache
	Get(ctx context.Context, key []byte) ([]byte, error)
	// Set sets the value for a given key in the cache with an optional expiration time.
	// If the key already exists in the cache, its value will be updated.
	// Returns nil if the key-value pair does not exist in the cache.
	Set(ctx context.Context, key, value []byte, expired int64) error

	// Incr increments the value for a given key in the cache.
	Incr(ctx context.Context, key []byte) error
}
