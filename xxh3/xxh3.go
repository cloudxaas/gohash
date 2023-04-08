package xxh3

import (
	"github.com/zeebo/xxh3"
)

// Sum64 gets the string and returns its uint64 hash value.
func Hash(key []byte) uint64 {
 	return xxh3.Hash(key)
}
