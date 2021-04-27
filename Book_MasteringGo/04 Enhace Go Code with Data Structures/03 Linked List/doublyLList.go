package main

import "fmt"

type NodeDLL struct {
	Value    int
	Previous *NodeDLL
	Next     *NodeDLL
}

var rootDLL = new(NodeDLL)

func addNodeDLL(t *NodeDLL, v int) int {
	if rootDLL == nil {
		t = &NodeDLL{v, nil, nil}
		rootDLL = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &NodeDLL{v, temp, nil}
		return -2
	}

	return addNodeDLL(t.Next, v)
}

func traverseDLL(t *NodeDLL) {
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

func reverse(t *NodeDLL) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func sizeDLL(t *NodeDLL) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupNodeDLL(t *NodeDLL, v int) bool {
	if rootDLL == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNodeDLL(t.Next, v)
}

func main() {
	fmt.Println(rootDLL)
	rootDLL = nil
	traverseDLL(rootDLL)
	addNodeDLL(rootDLL, 1)
	traverseDLL(rootDLL)
	addNodeDLL(rootDLL, 10)
	addNodeDLL(rootDLL, 5)
	addNodeDLL(rootDLL, 0)
	addNodeDLL(rootDLL, 0)
	traverseDLL(rootDLL)
	addNodeDLL(rootDLL, 100)
	fmt.Println("Size:", sizeDLL(rootDLL))
	traverseDLL(rootDLL)
	reverse(rootDLL)
}
