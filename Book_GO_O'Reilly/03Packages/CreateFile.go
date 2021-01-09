package main

import "os"

func main() {
	file, err := os.Create("Book_GO_O'Reilly/03Packages/new.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, _ = file.WriteString("Creando un achivo en Go")
}
