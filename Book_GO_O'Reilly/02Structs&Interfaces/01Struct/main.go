package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

//Funcion normal
/*
func circleArea(c *Circle) float64{
	return math.Pi*c.r*c.r
}

func main() {
	c := Circle{0,0,5}
	fmt.Println(circleArea(&c))
}

*/

//Funcion en forma de metodo

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

//Aca no necesitamos el operador &, Go ya lo sabe
func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(c.area())
	fmt.Println(r.area())
}
