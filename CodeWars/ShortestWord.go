package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "should test that the solution returns the correct value"
	words := strings.Split(s, " ")
	shortest := len(words[0])
	for _, word := range words {
		if len(word) < shortest {
			shortest = len(word)
		}
	}
	fmt.Println(shortest)
}
