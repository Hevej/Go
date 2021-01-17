package main

import "fmt"

func a() {
	fmt.Println("Inside a()")

	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()

	fmt.Println("About to call b()")
	b()
	//Aca se genera el panic() y el recover()
	//Esto nunca se alcanza:
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	//Esto nunca se alcanza:
	fmt.Println("Exiting b()")
}

func main() {
	/*
		panic():
			function that terminates the current flow of a Go program and starts panicking
		recover():
			function that allows you to take back control of a goroutine that jus panicked using panic()
	*/
	a()
	fmt.Println("main() ended!")
}
