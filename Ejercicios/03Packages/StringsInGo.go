package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	// Find the position into a bigger string
	fmt.Println(strings.Index("test", "e"))
	//Join a list of strings
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	//Repeat a string
	fmt.Println(strings.Repeat("a", 5))
	//Replace a string
	fmt.Println(strings.Replace("aaaa", "aa", "b", 2))
	//Split a string
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	//Lowercase letters
	fmt.Println(strings.ToLower("TEST"))
	//Uppercase letters
	fmt.Println(strings.ToUpper("test"))
	//Convert to a binary data and vice versa
	arr := []byte("test")
	fmt.Println(arr)
	str := string(arr)
	fmt.Println(str)
}
