package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	// This call returns a bufio.Scanner variable
	// Used with the Scan() function for reading from os.Stdin
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Stop reading data from standard input by pressing Ctrl+D
		fmt.Println(">", scanner.Text())
	}
}
