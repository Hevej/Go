package main

import "fmt"

type User interface {
	Permisos() int
	Nombre() string
}

type Admin struct {
	nombre string
}

type Editor struct {
	nombre string
}

func (admin Admin) Permisos() int {
	return 5
}

func (admin Admin) Nombre() string {
	return admin.nombre
}

func (editor Editor) Permisos() int {
	return 3
}

func (editor Editor) Nombre() string {
	return editor.nombre
}

func auth(user User) string {
	if user.Permisos() > 4 {
		return user.Nombre() + " tiene permisos de administrador"
	}
	return user.Nombre() + " tiene permisos de editor"
}

func main() {
	admin := Admin{"Luis"}
	editor := Editor{"Felipe"}

	usuarios := []User{admin, editor}

	for _, usuarios := range usuarios {
		fmt.Println(auth(usuarios))
	}
}
