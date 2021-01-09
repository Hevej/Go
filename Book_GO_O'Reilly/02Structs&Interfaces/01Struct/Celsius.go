package main

import (
	"fmt"
	"reflect"
)

type celsius float64

func (c celsius) String() string {
	return fmt.Sprintf("%.2f C", c)
}

type temperature struct {
	celsius
}

type AnotherTemp celsius

func main() {
	c := celsius(10.0)
	fmt.Println(c)

	// Hereda los metodos
	t := temperature{c}
	fmt.Println(t)

	// No hereda los metodos
	at := AnotherTemp(c)
	fmt.Println(at)
	fmt.Println(reflect.TypeOf(at))
}
