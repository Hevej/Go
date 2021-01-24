package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]

	//Allows to read a file all at once
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	//The value of each map key is of the type interface{},
	//which can be of any type
	var parsedData map[string]interface{}
	//Function used for putting the contents of the file into the parsedData map
	_ = json.Unmarshal([]byte(fileData), &parsedData)

	//Iterate over the map and get its contents
	for key, value := range parsedData {
		fmt.Println("key:", key, "value", value)
	}
}
