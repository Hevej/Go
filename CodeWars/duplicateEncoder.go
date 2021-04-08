package main

import (
	"fmt"
	"sort"
	"strings"
)

type sortWord interface {
	Rns []rune
}

func Contains(letters []int,letter int) int {
	sortedLetters := sort.Sort()
	sort.SearchInts(sortedLetters, letter)
	return int
}

func main() {
	word := "Success"
	rs := []rune(strings.ToLower(word))
	for _, letter := range rs {
		fmt.Println(Contains(rs, letter))
	}
	//( 0x28
	//) 0x29
}
