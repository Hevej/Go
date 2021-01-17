package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		// EL exit terminara un programa e imprimira en consola el estatus que yo elija (int)
		// Los defer nunca seran llamdos cuando exista un os.Exit
		os.Exit(1)
	}

	arguments := os.Args

	// El primer elemento del slice os.Args siempre sera el nombre del ejecutable
	fmt.Println(arguments[0])
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
