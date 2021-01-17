package main

import (
	"fmt"
	"runtime"
)

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}

	fmt.Println(structure)

	// Controla que se haga una go rutine, la del garbage collector
	runtime.GC()
	// Previene que el GC recolecte la variable structure muy temprano
	_ = structure[0]
}
