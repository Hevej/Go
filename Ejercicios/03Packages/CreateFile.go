package main

import "os"

func main() {
	file, err := os.Create("Ejercicios/03Packages/new.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, _ = file.WriteString("Creando un achivo en Go")
}
