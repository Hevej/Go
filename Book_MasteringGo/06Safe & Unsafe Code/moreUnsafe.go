package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0, 1, -2, 3, 4}
	pointer := &array[0]
	//Apunta al cero
	fmt.Println(*pointer, " ")
	/* Se convierte a un puntero unsafe
	 Despues se convierte a uintptr
		- UintPtr es un tipo de int sin signo que es capaz de almacenar data de un puntero
		- Significa que es del mismo tamaño que un puntero
	Esto se le suma el espacio que ocupa el elemento 0 y se guarda todo en la
	variable memoryAddress
	*/
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	/*
		unsafe-Sizeof(array[0]) me da la memoria que ocupa cada elemento del array,
		asi que este valor es añadido a cada iteración para continuar con el proximo elemento
	*/
	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
	fmt.Println()
	//Se accede a un elemento en el array que no existe usando punteros y direcciones
	// de memoria
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	//El problema esta en que no se produce un error y el programa devuelve
	//un numero aleatorio.
	fmt.Print("One more: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println()

}
