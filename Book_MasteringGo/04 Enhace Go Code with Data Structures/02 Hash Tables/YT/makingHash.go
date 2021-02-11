package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket is a linked list in each slot of the hash table array
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	//Crea una hashtable
	//Lo localizo por que tengo que devolver un puntero
	result := &HashTable{}
	//Navego entre la estructura interna de result, un array de longitud 7
	for i := range result.array {
		//A cada elemento le asigno un puntero de bucket{}, el tipo que espera
		//Al iniciarlizarlo asi, quedara cada bucket guardando una dirección y no
		//quedara el array nulo
		result.array[i] = &bucket{}
	}
	//Devuelvo la hash table ya hecha
	return result
}

func main() {
	//Retorna una hash table donde cada slot tiene un bucket
	testHashTable := Init()
	fmt.Println(testHashTable)
	fmt.Println("Estado incompleto hasta aprender recievers")
}
