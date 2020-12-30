package main

import (
	"fmt"
	"sort"
)

func greatest(args ...float64) float64  {
	sort.Float64s(args)
	return args[len(args)-1]
}

func main() {
	s :=[]float64{1.24,4.0,5,6.82,1.28,4.84,9.72,0.72,2.34}
	fmt.Println(greatest(s ...))
}
