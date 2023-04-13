package city64custom

const (
	kMul = 0x9ddfea08eb382d69
	kSeed = 0xc70f6907b2205f5a
)

func fetch64(p []byte) uint64 {
	return uint64(p[0]) |
		uint64(p[1])<<8 |
		uint64(p[2])<<16 |
		uint64(p[3])<<24 |
		uint64(p[4])<<32 |
		uint64(p[5])<<40 |
		uint64(p[6])<<48 |
		uint64(p[7])<<56
}

func fetch32(p []byte) uint32 {
	return uint32(p[0]) |
		uint32(p[1])<<8 |
		uint32(p[2])<<16 |
		uint32(p[3])<<24
}

func rotate(val uint64, shift int) uint64 {
	if shift == 0 || shift == 64 {
	    return val
	} 
	return (val >> shift) | (val << (64 - shift))
}

func hash128to64(x [2]uint64) uint64 {
	const (
		c1 = 0xff51afd7ed558ccd
		c2 = 0xc4ceb9fe1a85ec53
	)

	h := x[0] ^ x[1]
	h = (h ^ (h >> 33)) * c1
	h = (h ^ (h >> 33)) * c2
	h ^= (h >> 33)
	return h
}

func Hash(data []byte) uint64 {
	length := uint64(len(data))
	p := 0
	limit := len(data)
	h := kSeed ^ (length * kMul)

	for p+8 <= limit {
		w := fetch64(data[p:])
		h += w
		h = rotate(h, 29) * kMul
		p += 8
	}
	if p+4 <= limit {
		w := fetch32(data[p:])
		h += uint64(w)
		h *= kMul
		p += 4
	}
	if p < limit {
		w := uint64(data[p])
		h += w
		h *= kMul
	}
	return hash128to64([2]uint64{h, kMul})
}
