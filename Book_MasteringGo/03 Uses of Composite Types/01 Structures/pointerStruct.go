package main

type myStructure struct {
	Name    string
	Surname string
	Height  uint32
}

func createStruct(n, s string, h uint32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}

func main() {

}
