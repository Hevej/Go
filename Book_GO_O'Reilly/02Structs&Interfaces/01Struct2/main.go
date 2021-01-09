package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (u User) nombreCompleto() string {
	return u.nombre + " " + u.apellido
}

//Pasando el puntero me permite mutar la estructura
func (u *User) setName(n string) {
	u.nombre = n
}

func main() {
	//Me crea una instancia de User y en forma de puntero
	luis := new(User)
	luis.nombre = "Luis Felipe"
	luis.apellido = "Pinto Ortegón"
	luis.edad = 25
	fmt.Println(luis)

	luis.setName("Aaron")
	//Ahora hacemos uso del metodo
	fmt.Println(luis.nombreCompleto())
}
