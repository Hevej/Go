package main

import "fmt"

type XYZ struct {
	X int
	Y int
	Z int
}

func main() {
	var s1 XYZ
	//I don´t need to define an inital value for every field of the structure
	fmt.Println(s1.Y, s1.Z)

	p1 := XYZ{23, 12, -2}
	//Structure literal
	p2 := XYZ{Z: 12, Y: 13}
	fmt.Println(p1)
	fmt.Println(p2)

	pSlice := make([]XYZ, 4)
	pSlice[2] = p1
	pSlice[0] = p2
	fmt.Println(pSlice)
	//Changing the value of the original structure will
	//have no effect on the objects of the array
	p2 = XYZ{1, 2, 3}
	fmt.Println(pSlice)
}
