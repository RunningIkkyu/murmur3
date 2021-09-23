package murmur3_test

import (
	"murmur3"
	"testing"
)

type testdata struct {
	key  string
	seed uint32
	hash uint32
}

// Online validate: http://murmurhash.shorelabs.com
var td = []testdata{
	{key: "hello", seed: 1, hash: 3142237357},
	{key: "world", seed: 1, hash: 1648897759},
	{key: "key1", seed: 1, hash: 1833321129},
	{key: "key1", seed: 2, hash: 386463789},
	{key: "key1", seed: 3, hash: 201916224},
	{key: "key2", seed: 1, hash: 2866059576},
	{key: "key2", seed: 2, hash: 4257591023},
	{key: "key2", seed: 3, hash: 2337524217},
	{key: "key3", seed: 1, hash: 1632542396},
	{key: "key3", seed: 2, hash: 1921854502},
	{key: "key3", seed: 3, hash: 245936421},
}

func TestMurmur3(t *testing.T) {
	for _, d := range td {
		expected := murmur3.Murmur32([]byte(d.key), int(d.seed))
		if expected != d.hash {
			t.Logf("testdata: %+v, expected: %d, got: %d", d, expected, d.hash)
			t.Fail()
		}
	}
}
