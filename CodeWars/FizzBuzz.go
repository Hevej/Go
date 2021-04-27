package main

import (
	"fmt"
	"math"
)

func main() {
	n := make([]int, 10)
	for i := 1; i <= 10; i++ {
		n[i-1] = i
	}
	fmt.Println(n)
	//n := []int{1,2,3,4,5,6,7,8,9,10}
	for _, numero := range n {
		div1 := math.Mod(float64(numero), 3.0) == 0.0
		div2 := math.Mod(float64(numero), 5.0) == 0
		if div1 && div2 {
			fmt.Println("FizzBuzz")
		} else if div1 && !div2 {
			fmt.Println("Fizz")
		} else if !div1 && div2 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(numero)
		}
	}
}
