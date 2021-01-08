package main

import "fmt"

func parImpar(numero int) (int, bool) {
	val := numero / 2
	return val, val%2 != 0
}

func main() {
	var numero int
	_, _ = fmt.Scanf("%d", &numero)
	res1, res2 := parImpar(numero)
	fmt.Println(res1, res2)
}
