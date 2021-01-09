package main

import (
	"fmt"
	"os"
)

func main() {
	//Open current directory
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	//If n<=0 Readdir returns all the file info
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
