package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice1 = append(slice1, 4, 5)
	slice2 := make([]int, 2)
	fmt.Println(slice1, slice2)
}
