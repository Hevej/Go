package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Calculo interface {
	Interes() float64
	Final() float64
}

type Compuesto struct {
	presente float64
	periodos int
	tasa     float64
}

type Simple struct {
	presente float64
	periodos int
	tasa     float64
}

func (s Simple) Interes() float64 {
	i := s.presente * float64(s.periodos) * s.tasa
	return i
}

func (s Simple) Final() float64 {
	return s.Interes() + s.presente
}

func (c Compuesto) Interes() float64 {
	i := c.presente * (math.Pow(1+c.tasa, float64(c.periodos)) - 1)
	return i
}

func (c Compuesto) Final() float64 {
	return c.Interes() + c.presente
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Ingrese argumentos.")
		os.Exit(1)
	}

	args := os.Args
	vP, _ := strconv.ParseFloat(args[2], 64)
	n, _ := strconv.Atoi(args[3])
	tS, _ := strconv.ParseFloat(args[4], 64)
	argumentos := []Calculo{}

	switch args[1] {
	case "s":
		argumentos = append(argumentos, Simple{vP, n, tS})
		break
	case "c":
		argumentos = append(argumentos, Compuesto{vP, n, tS})
		break
	default:
		break
	}

	for _, usuarios := range argumentos {
		fmt.Println("Intereses: ", Calculo(usuarios).Interes(), " Total: ", Calculo(usuarios).Final())
	}

	/*
		sim := Simple{1000, 3, 0.05}
		com := Compuesto{1000, 3, 0.05}
		fmt.Println(Calculo(sim).Interes())
		fmt.Println(Calculo(com).Interes())
	*/
}
