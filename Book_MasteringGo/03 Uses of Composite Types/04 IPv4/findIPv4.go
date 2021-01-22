package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

//Contains the definition of the reguñar expression that will
//help me to discover an IPv4 address inside a function
func findIP(input string) string {
	//Buscas las posibles combinaciones para una triada de 0-255 de una IP
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	//The think that we are looking for has four distinct parts, and each
	//one of them must match partIP
	//Busca toda la IP completa
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s logFile \n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file %s\n", err)
			os.Exit(-1)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}

			//Llamo a la función pora buscar la IP en el texto
			ip := findIP(line)
			//Double-check that we are dealing whit a valid IPv4
			//Si no es correcta, devuelve nil
			trial := net.ParseIP(ip)
			//Si la conversion no es correcta, devuelve nil
			if trial.To4() == nil {
				continue
			}
			fmt.Println(ip)
		}
	}
}
