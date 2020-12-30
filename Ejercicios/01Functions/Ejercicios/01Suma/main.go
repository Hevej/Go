package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(sli ...int) (total int){

	for _,v:=range sli{
		total += v
	}
	return
}

func main() {
	//Con un slice
	/*
	xs :=[]int{1,2,3,4,5}
	fmt.Println(sum(xs...))*/

	var conjuntoNumeros string
	fmt.Println("Escriba los numeros separados por espacios")
	//Se crea el scanner para escuchar un stream
	scanner := bufio.NewScanner(os.Stdin)
	//Se lee el stream
	scanner.Scan()
	//Se obtienen los datos
	conjuntoNumeros = scanner.Text()
	//Se pasan a un slice de strings
	sliceNumbers := strings.Fields(conjuntoNumeros)
	//convierto de Str a Int con Atoi
	var V []int
	for _, v:=range sliceNumbers{
		//fmt.Println(reflect.TypeOf(v),v)
		num, _ := strconv.Atoi(v)
		V = append(V,num)
	}

	//fmt.Println(conjuntoNumeros)
	//fmt.Println(sliceNumbers)
	//fmt.Println(reflect.TypeOf(V),V)

	fmt.Println(sum(V ...))
}
