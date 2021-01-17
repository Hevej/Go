package main

// Map that stores plain values without pointers

import (
	"fmt"
	"runtime"
)

func main() {
	var N = 40
	myMap := make(map[int]int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = value
	}
	fmt.Println(myMap)
	runtime.GC()
	_ = myMap[0]
}
