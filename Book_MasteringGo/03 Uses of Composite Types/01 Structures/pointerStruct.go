package main

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	Height  uint32
}

//You are allowed to check whether the provided information is both correct and valid
func createStruct(n, s string, h uint32) *myStructure {
	if h > 300 {
		h = 0
	}
	//It is perfectly leagl for a Go function to return the memory address of a local variable
	return &myStructure{n, s, h}
}

//Both functions works fine, but are different in the way they work.
func retStructure(n, s string, h uint32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

func main() {
	s1 := createStruct("Mihalis", "Tsoukalos", 123)
	s2 := retStructure("Mihalis", "Tsoukalos", 123)
	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)
}
