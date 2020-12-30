package main

import "fmt"

func swap(nInt1 *int, nInt2 *int)   {
	aux := *nInt1
	*nInt1 = *nInt2
	*nInt2 = aux
}

func main() {
	num1 := 1
	num2 := 2
	swap(&num1,&num2)
	fmt.Println(num1,num2)
}
