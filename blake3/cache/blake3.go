package blake3cache

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/valyala/bytebufferpool"
	"github.com/zeebo/blake3"
)

var cache *fastcache.Cache

func New(cacheSize int) {
	cache = fastcache.New(cacheSize)
}

func Hash(key []byte) *bytebufferpool.ByteBuffer {
	//if cache == nil {
	//	panic("Cache is not initialized. Call InitCache() first.")
	//}

	cachedValue := cache.Get(nil, key)
	if cachedValue != nil {
		buf := bytebufferpool.Get()
		buf.Write(cachedValue)
		return buf
	}

	hasher := blake3.New()
	hasher.Write(key)
	hash := hasher.Sum(nil)

	newValue := bytebufferpool.Get()
	newValue.Write(hash)

	cache.Set(key, newValue.B)

	return newValue
}

func Release(buf *bytebufferpool.ByteBuffer) {
	bytebufferpool.Put(buf)
}