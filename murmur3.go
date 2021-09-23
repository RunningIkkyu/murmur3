package murmur3

import (
	"math/bits"
)

const (
	DefaultSeed = 0x01
)

func Murmur32(key []byte, seeds ...int) (hash uint32) {
	seed := DefaultSeed
	if len(seeds) > 0 {
		seed = seeds[0]
	}
	length := len(key)
	hash = uint32(seed)
	var c1 uint32 = 0xcc9e2d51
	var c2 uint32 = 0x1b873593
	r1 := 15
	r2 := 13
	var m uint32 = 5
	var n uint32 = 0xe6546b64

	buf := NewByteReader(key)
	for {
		k, err := buf.ReadUint32()
		if err != nil {
			break
		}

		k = k * c1
		k = bits.RotateLeft32(k, r1)
		k = k * c2

		hash = hash ^ k
		hash = bits.RotateLeft32(hash, r2)
		hash = hash*m + n
	}

	// Remain bytes
	t, err := buf.ReadUint32Anyway()
	if err == nil {
		//fmt.Printf("buf is: %v, t is %d\n", buf.buf, t)

		t = t * c1
		t = bits.RotateLeft32(t, r1)
		t = t * c2

		hash = hash ^ t
	}

	hash = hash ^ uint32(length)

	hash = hash ^ (hash >> 16)
	hash = hash * 0x85ebca6b
	hash = hash ^ (hash >> 13)
	hash = hash * 0xc2b2ae35
	hash = hash ^ (hash >> 16)

	return hash
}
