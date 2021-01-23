package main

import (
	"encoding/json"
	"fmt"
)

type R struct {
	Name    string
	Surname string
	Tel     []T
}

type T struct {
	Mobile bool
	Number string
}

func main() {
	myRecord := R{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []T{
			T{Mobile: true, Number: "1234-567"},
			T{Mobile: true, Number: "1234-abcd"},
			T{Mobile: false, Number: "abcc-567"},
		},
	}

	//Accepts a reference to the myRecord variable
	//Converts into a JSON format
	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rec))

	var unRec R
	//Gets JSON and converts it into a Go structure
	//Requieres a pointer argument
	err1 := json.Unmarshal(rec, &unRec)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(unRec)
	fmt.Println(unRec.Tel)
}
