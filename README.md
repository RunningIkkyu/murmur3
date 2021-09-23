# Murmur3

A go implementation for Murmur3 Hash.

The algorithm description can be found in Wiki: 

You could also test murmur3 hash online here:

http://murmurhash.shorelabs.com


# Quick Start

Only need one function call:

```go
key := []byte("Murmur3")
seed := 0x01
hash := Murmur3(key, seed)
fmt.Printf("Murmur3 hash of key: %s is %s.\n", key, hash)
```
