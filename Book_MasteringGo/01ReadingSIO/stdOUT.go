package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me on argument!"
	} else {
		myString = arguments[1]
	}

	// No se puede diferenciar la data de Stdout y Stderr en consola
	io.WriteString(os.Stdout, "This is Standard ouput\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
