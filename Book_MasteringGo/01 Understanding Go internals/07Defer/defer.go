package main

import "fmt"

func d3() {
	// Esta es la forma valida de usar defer y la variable i en un ciclor for
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	fmt.Println()
}

func d2() {
	// Cuando el ciclo for termina, el valor de i es cero
	// Pero la funcion en el defer se ejecuta despues de que el for acabe
	// Lo que significa que se ejecuta 3 veces en cero
	// Es una forma invalidad de usar la variable i
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d1() {
	// Los defer son LIFO (last in, first out)
	// Se llamara primero el ultimo defer
	// asi mismo, se llamara de ultimo el primer defer
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func main() {
	d1()
	d2()
	d3()
}
