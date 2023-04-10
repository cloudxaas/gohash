package xxh3custom

func avalanche(h64 uint64) uint64 {
	h64 ^= h64 >> 33
	h64 *= 0xff51afd7ed558ccd
	h64 ^= h64 >> 33
	h64 *= 0xc4ceb9fe1a85ec53
	h64 ^= h64 >> 33
	return h64
}

func mix2accs(acc uint64, key uint64) uint64 {
	val := acc + key
	return val * 0x9fb21c651e98df25
}

func len1to3_64b(data []byte, length int, seed uint64) uint64 {
	c1 := uint64(data[0]) * 0x87c37b91114253d5
	c2 := uint64(data[length>>1]) * 0x4cf5ad432745937f
	c3 := uint64(data[length-1]) * 0x52dce729
	h64 := avalanche(c1 ^ c2 ^ c3 ^ seed)
	return h64
}

func readUint64LE(data []byte) uint64 {
	return uint64(data[0]) |
		uint64(data[1])<<8 |
		uint64(data[2])<<16 |
		uint64(data[3])<<24 |
		uint64(data[4])<<32 |
		uint64(data[5])<<40 |
		uint64(data[6])<<48 |
		uint64(data[7])<<56
}
func Hash(data []byte) uint64 {
	return Hash64(data, len(data), uint64(0))
}
func Hash64(data []byte, length int, seed uint64) uint64 {
	acc := [2]uint64{seed, seed}

	if length <= 3 {
		return len1to3_64b(data, length, seed)
	}

	for i := 0; i+16 <= length; i += 16 {
		acc[0] = mix2accs(readUint64LE(data[i:i+8]), acc[0])
		acc[1] = mix2accs(readUint64LE(data[i+8:i+16]), acc[1])
	}

	length &= 15
	if length != 0 {
		acc[0] = mix2accs(readUint64LE(data[length-8:length]), acc[0])
		acc[1] = mix2accs(readUint64LE(data[length&^8:length]), acc[1])
	}

	h64 := acc[0] ^ avalanche(acc[1])
	return h64
}
