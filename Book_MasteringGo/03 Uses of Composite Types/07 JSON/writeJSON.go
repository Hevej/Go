package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Re struct {
	Name    string
	Surname string
	Tel     []Tel
}

type Tel struct {
	Mobile bool
	Number string
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Re{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Tel{
			Tel{Mobile: true, Number: "1234567"},
			Tel{Mobile: true, Number: "1234-abcd"},
			Tel{Mobile: false, Number: "abcc-567"},
		},
	}
	//By using os.Stdout, the data will be printed on the screen instead of being
	//saved into a file
	saveToJSON(os.Stdout, myRecord)
}
