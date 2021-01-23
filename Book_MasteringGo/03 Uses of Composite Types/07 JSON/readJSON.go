package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

//Decoding a JSON according to a data structure
func loadJSON(filename string, key interface{}) error {
	//Open the file
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	//Create a JSON decode variable
	decodeJson := json.NewDecoder(in)
	//Actually deconding the contents of the file
	//Stores the value in pointer key
	err = decodeJson.Decode(key)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]

	//Interfaz con la estructura deseada
	var myRecord Record
	//Decodifica y guarda la información en myRecord
	err := loadJSON(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}
}
