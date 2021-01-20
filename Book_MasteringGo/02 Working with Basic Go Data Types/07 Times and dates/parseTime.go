package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myTime string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	myTime = os.Args[1]

	//Parse and hour and minute string whit "15:04"
	//More const in https://golang.org/src/time/format.go
	d, err := time.Parse("15:04", myTime)
	//The value of the err variable tells you wheteher the parsing was successful or not
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	} else {
		fmt.Println(err)
	}
}
