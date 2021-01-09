package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonalSort interface {
	Len() int
	Les(i, j int) bool
	Swap(i, j int)
}

type ByName []Person
type ByAge []Person

// Retorna la longitud del slice
func (ps ByName) Len() int {
	return len(ps)
}

// Retorna una comparación, entre items de posiciones i, j
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

// Cambia de posición los items
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ag ByAge) Len() int {
	return len(ag)
}

func (ag ByAge) Less(i, j int) bool {
	return ag[i].Age < ag[j].Age
}

func (ag ByAge) Swap(i, j int) {
	ag[i], ag[j] = ag[j], ag[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	//fmt.Println(reflect.TypeOf(ByName(kids)))
	v := ByName(kids)
	// Ordena nuestra lista de personas
	sort.Sort(v)
	fmt.Println(kids)

	// Ordenamiento propio
	v.Swap(0, 1)
	fmt.Println(v)
	fmt.Println("De longitud", v.Len())
	fmt.Println(v[0].Name, " < ", v[1].Name, " = ", v.Less(0, 1))

	// Ahora lo puedo hacer con un objeto ByAge
	v2 := ByAge(kids)

	sort.Sort(v2)

	// Ordenamiento propio
	v.Swap(0, 1)
	fmt.Println(2)
	fmt.Println("De longitud", v2.Len())
	fmt.Println(v2[0].Name, " < ", v2[1].Name, " = ", v2.Less(0, 1))
}
