package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// remember to always close opened files
	defer f.Close()

	// create a hasher
	h := crc32.NewIEEE()
	// copy the file into the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, err = io.Copy(h, f)
	// we don't care about how many bytes were written, but we do want to
	// handle the error
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("Book_GO_O'Reilly/03Packages/test1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("Book_GO_O'Reilly/03Packages/test2.txt")
	if err != nil {
		return
	}

	// Compara ambos archivos por hash con crc32, si ambos dan igual
	// es altamente probable que sean el mismo archivo
	// si no dan igual, son diferentes.
	fmt.Println(h1, h2, h1 == h2)
}
