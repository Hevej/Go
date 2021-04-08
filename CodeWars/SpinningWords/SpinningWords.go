package main

import (
	"fmt"
	"os"
	s "strings"
)

func SpinWords(str string) string {
	//Volver a runa, string inmutable
	r := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	palabras := os.Args[1:]
	for i, elem := range palabras {
		if len(elem) > 4 {
			palabras[i] = SpinWords(elem)
		}
	}
	fmt.Println(s.Join(palabras, " "))
}

/*
package kata

import (
	s "strings"
)

func Reverse(str string) string {
  r := []rune(str)
	for i,j := 0, len(str)-1; i<j ; i,j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func SpinWords(str string) string  {
  palabra := s.Split(str, " ")
	for i, elem := range palabra {
		if len(elem)>4 {
			palabra[i] = Reverse(elem)
		}
	}
	return s.Join(palabra," ")
}
*/
