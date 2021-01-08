package main

import (
	"container/list"
	"fmt"
)

func main() {
	//Create a linked list
	var x list.List
	//Values are appended to the list
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	x.PushBack(4)

	fmt.Println(x.Front())

	//Loop over each item in the list by getting the first element
	//then, follow all the links until we reach nil
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
