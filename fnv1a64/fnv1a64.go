package fnv1a64

const (
    offset64 = 14695981039346656037
    prime64  = 1099511628211
)

func Hash(data []byte) uint64 {
    var h uint64 = offset64
    for _, b := range data {
        h ^= uint64(b)
        h *= prime64
    }
    return h
}
