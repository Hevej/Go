package main

import (
	"fmt"
	"os"
	"strconv"
)

func Odss(n int) int {
	odd := -1
	for i := 0; i < n; i++ {
		odd = odd + 2
	}
	return odd
}

func main() {
	args := os.Args
	n, _ := strconv.Atoi(args[1])
	max := 0
	r := 0
	for i := 0; i <= n; i++ {
		max += i
	}
	min := max - n + 1
	for i := Odss(min); i <= Odss(max); i = i + 2 {
		r += i
	}
	fmt.Println(Odss(min), Odss(max), r)
}
