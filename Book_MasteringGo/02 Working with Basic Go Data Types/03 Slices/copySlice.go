package main

import "fmt"

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	//copy(destination, source)
	//The last two elements of a a6 will remain the same
	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)

	//Only the first four elements of b6 will be copied to b4
	copy(b4, b6)
	fmt.Println("b6:", b6)
	fmt.Println("b4:", b4)
	fmt.Println()

	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	//Arrays is converted to a slice with the help of the [:] notation
	copy(s6, array4[0:])
	fmt.Println("array4:", array4[0:])
	fmt.Println("s6:", s6)
	fmt.Println()

	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, -7, 7, -7, 7}
	//If I try to copy an array into a slice or vice versa without using the [:] notation
	//the program will fail to compile
	copy(array5[0:], s7)
	fmt.Println("array5:", array5)
	fmt.Println("s7:", s7)
}
