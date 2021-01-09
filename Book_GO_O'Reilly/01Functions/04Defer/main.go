package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	//Programa a la funcion para ser llamada
	//cuando la funcion que lo contiene, complete.
	defer second()
	first()
}
