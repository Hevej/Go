package main

import (
	"fmt"
	"hash/crc32"
)

// Non-cryptographic hash functions:
// has package includes:
// adler32, crc32, crc64 and fnv

func main() {
	// create a hasher
	h := crc32.NewIEEE()
	fmt.Println(h)
	// write our data to it
	_, _ = h.Write([]byte("test"))
	// calculate the crc32 checksum
	fmt.Println(h)
	v := h.Sum32()
	fmt.Println(v)
}
