package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	integers := make([]int, 2)
	fmt.Println(integers)
	//Elimino todos los elementos del slice
	integers = nil
	fmt.Println(integers, cap(integers))

	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:]
	fmt.Println(anArray)
	fmt.Println(refAnArray)
	//El slice apunta hacia el mismo espacio de memoria que el array
	anArray[4] = -100
	fmt.Println(refAnArray)

	//One dimension slice
	s := make([]byte, 5)
	fmt.Println(s)
	//Two dimensions slice
	//The zero value for the slice type is nil
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()

	//Agrego valores al slice
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
	}

	//Leo esos valores del slice 2D
	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value:", y)
		}
	}
}
