//Correrlo con: go run selectColumn.go 2 logs.txt
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	//Necesito al menos dos argumentos, columna y ruta de archivo
	//Solo un numero de columna para multoples archivos
	if len(arguments) < 3 {
		fmt.Println("Usage: selectColumn column <file1> [<file2>" +
			" [...<fileN>]]")
		os.Exit(1)
	}

	//Convierte el string de la columna a un int
	column, err := strconv.Atoi(arguments[1])
	//Si no es un numero, devuelve un mensaje de error y termina con el programa
	if err != nil {
		fmt.Println("Column value is not an integer:", column)
		return
	}

	//Si el numero de la columna es negativo, invalida y termina el programa
	if column < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}

	//Se usa un for para navegar entre archivos, que se pasan como argumentos
	for _, filename := range arguments[2:] {
		//Nombre de la ruta
		fmt.Println("\t\t", filename)
		//Abre el archivo
		open, err := os.Open(filename)
		//Si hay error manda el mensaje de error y continua al otro archivo
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		//Defer para cerrar archivos al final
		defer open.Close()

		//Retorna un Reader de buffer predeterminado = 4096Bytes
		r := bufio.NewReader(open)
		//While true
		for {
			//Lee cada linea hasta el salto de linea
			//Retorna un byte slice
			line, err := r.ReadString('\n')

			//EOF = NOT MORE INPUT AVAILABLE
			//Cuando ya no hay nada en el buffer, rompe el while
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error reading file %s\n", err)
			}

			//Splits a string based on the whitespace and return a slice of strings
			data := strings.Fields(line)
			fmt.Println(data)
			//Si la columna existe, devuelve la columna
			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}
	}
}
