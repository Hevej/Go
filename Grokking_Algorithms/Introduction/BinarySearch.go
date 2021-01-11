package main

import "fmt"

func binarySearch(list []int, number int) (n bool) {
	min := 0
	max := len(list) - 1

	for min <= max {
		middle := (max + min) / 2
		if number == list[middle] {
			n = true
			break
		} else if list[middle] > number {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return
}

func main() {
	myList := []int{1, 3, 5, 7, 9}
	var numero int
	fmt.Println("Ingrese un numero")
	fmt.Scanf("%d", &numero)
	fmt.Println(binarySearch(myList, numero))
}
