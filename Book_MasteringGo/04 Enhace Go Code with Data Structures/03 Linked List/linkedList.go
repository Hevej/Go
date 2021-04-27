package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

//New returns a pointer (*Node) to a newly allocated zero value of type Node
//Root holds the first element of the linked list
var root = new(Node)

//Add new nodes to the linked list
func addNode(t *Node, v int) int {
	//If we are dealing with an empty linked list, it is created
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}

	//Checking if the value that I want to add is already in the list, in this position
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	//I check whether I have reached the end of the linked list, then, add a node
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	//Repeat the same process
	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// Give an element and checks if it exits into the list
func lookupNode(t *Node, v int) bool {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

// Returns the size of the linked list
func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, -1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)
	addNode(root, 100)
	traverse(root)

	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	if lookupNode(root, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
}
