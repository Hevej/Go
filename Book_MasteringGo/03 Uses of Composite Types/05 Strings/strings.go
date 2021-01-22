package main

import "fmt"

func main() {
	//Each \xAB sequence represents a single character of sLiteral
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	//Return AB part of \xAB
	fmt.Printf("x: %x\n", sLiteral)
	//Return the number of characters of sLiteral
	fmt.Printf("sLiteral length: %d\n", len(sLiteral))

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	fmt.Println()

	//String that is safely escaped
	fmt.Printf("q: %q\n", sLiteral)
	//ASCII-only output
	fmt.Printf("+q: %+q\n", sLiteral)
	//Spaces printed between bytes
	fmt.Printf("x: % x\n", sLiteral)

	s2 := "€£³"
	for x, y := range s2 {
		//Print the characters in the U+0058 format
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}
	//Unicode characters that they are a byte size larger than the number
	//of characters in it
	fmt.Printf("s2 length: %d\n", len(s2))

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
