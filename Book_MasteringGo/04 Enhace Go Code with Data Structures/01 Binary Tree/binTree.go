package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Definition of a node
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

//The traverse() function reveals how you can visit all of the nodes
//of a binary tree using recursion
func traverse(t *Tree) {
	//Si no hay nada salgo de la recursion
	if t == nil {
		return
	}
	//Navega el nodo izquierdo
	traverse(t.Left)
	//Imprime el nodo padre y deja un espacio
	fmt.Print(t.Value, " ")
	//Navega el nodo derecho
	traverse(t.Right)
}

//The create() function is only used for populating the binary tree with random integers
func create(n int) *Tree {
	var t *Tree
	//Used for populating the tree with random numbers, as we do not have any real data
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		//retorna un uint de [0,2*n)
		temp := rand.Intn(n * 2)
		//Al arbol, le inserta el numero numero random
		t = insert(t, temp)
	}
	//Arbol ya lleno
	return t
}

func insert(t *Tree, v int) *Tree {
	//Si el arbol esta vacio, le asigna valor al root
	if t == nil {
		return &Tree{nil, v, nil}
	}
	//Verifica si el valor ya existe, devolviendo el arbol
	if v == t.Value {
		return t
	}
	//Si es menor al valor del nodo, lo ingresa a la izquierda
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	//Si es menor al valor del nodo, lo ingresa a la derecha
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}
