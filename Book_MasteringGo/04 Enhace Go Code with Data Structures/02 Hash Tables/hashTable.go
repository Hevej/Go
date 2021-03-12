package main

import "fmt"

// Definition of the node
const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return i % size
}

func insert(hash *HashTable, value int) int {
	//Crea un identificador para el valor en la hashtable
	index := hashFunction(value, hash.Size)
	//Crea un nodo con el valor
	element := Node{Value: value, Next: hash.Table[index]}
	//Lo adiciona a la hash table
	hash.Table[index] = &element
	return index
}

//Solo imprime la tabla
func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

//Determina en donde esta un elemento
func lookup(hash *HashTable, value int) bool {
	//Obtengo el identificador
	index := hashFunction(value, hash.Size)
	//Si el identificador tiene valor
	if hash.Table[index] != nil {
		//Obtengo la linked list
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			//Aca me va cambiando los valores para revisar el siguiente
			t = t.Next
		}
	}
	return false
}

func main() {
	//Creo la estructura - backbone
	table := make(map[int]*Node, SIZE)
	//Creo la hastable y retorno un address
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces:", hash.Size)
	//Inserto valores del 0 al 119
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}
	traverse(hash)


	fmt.Println(lookup(hash, 20))
	fmt.Println(lookup(hash, 220))

}
