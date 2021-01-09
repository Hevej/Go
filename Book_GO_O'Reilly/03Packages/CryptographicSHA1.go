package main

import (
	"crypto/sha1"
	"fmt"
	"reflect"
)

func main() {
	// computes a 160-bit hash
	h := sha1.New()
	fmt.Println(reflect.TypeOf(h))
	h.Write([]byte("test"))
	fmt.Println(h)
	// appends the current hash to []byte{}
	// and returns the resulting slice
	bs := h.Sum([]byte{})
	// There is no native type to represent a 160-bit number
	// so we use a slice of 20 bytes indstead
	fmt.Println(bs)
}
