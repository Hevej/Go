package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number1, _ := strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	number2, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	text1 := scanner.Text()

	fmt.Println(number1 + i)
	fmt.Printf("%.1f\n", number2+d)
	fmt.Println(s + text1)

}
