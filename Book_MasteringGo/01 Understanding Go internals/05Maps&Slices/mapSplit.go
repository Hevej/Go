package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// Split a map iunto a map of maps : sharding

func main() {
	var N = 400
	split := make([]map[int]int, 200)

	for i := range split {
		split[i] = make(map[int]int)
	}

	for i := 0; i < N; i++ {
		value := int(i)
		split[i%200][value] = value
		fmt.Println(split[i%200])
	}

	fmt.Println(reflect.TypeOf(split))
	runtime.GC()
	_ = split[0][0]
}
