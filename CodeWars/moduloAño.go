package main

import (
	"fmt"
	"math"
)

func main() {
	year := float64(2000)
	rest := math.Mod(year, 100)
	fmt.Println(int(year/100 + math.Ceil(rest/100)))
}
