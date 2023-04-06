package blake3cache

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/zeebo/blake3"
)

type Cache struct {
	cache *fastcache.Cache
}

func New(size int) *Cache {
	return &Cache{
		cache: fastcache.New(size),
	}
}

func (c *Cache) Hash(key []byte) []byte {
	// Check if the key exists in the cache
	v := c.cache.Get(nil, key)
	if len(v) > 0 {
		return v
	}

	// Compute the Blake3 hash of the key
	hash := blake3.Sum256(key)

	// Add the hash to the cache
	c.cache.Set(key, hash[:])

	return hash[:]
}

func (c *Cache) HashBig(key []byte) []byte {
	// Check if the key exists in the cache
	v := c.cache.GetBig(nil, key)
	if len(v) > 0 {
		return v
	}

	// Compute the Blake3 hash of the key
	hash := blake3.Sum256(key)

	// Add the hash to the cache
	c.cache.SetBig(key, hash[:])

	return hash[:]
}
