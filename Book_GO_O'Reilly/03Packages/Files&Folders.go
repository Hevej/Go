package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*
		file, err := os.Open("Book_GO_O'Reilly/03Packages/test.txt")
		if err != nil {
			//handle the error
			return
		}
		defer file.Close()

		//get the file size
		stat, err := file.Stat()
		if err != nil {
			return
		}

		//read the file
		bs := make([]byte, stat.Size())
		_, err = file.Read(bs)
		if err != nil{
			return
		}

		str := string(bs)
		fmt.Println(str)
	*/
	bs, err := ioutil.ReadFile("Book_GO_O'Reilly/03Packages/test.txt")
	if err != nil {
		panic(err)
	}
	str := string(bs)
	fmt.Println(str)
}
