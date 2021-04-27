package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "pp"
	rs := strings.ToLower(word)
	fmt.Println(rs)
	for _, letter := range rs {
		count := strings.Count(rs, string(letter))
		if count == 1 {
			rs = strings.Replace(rs, string(letter), "(", 1)
		} else {
			rs = strings.Replace(rs, string(letter), ")", -1)
		}
		if count == len(rs) {
			if rs == "))))))" {
				rs = ")))))("
			}
			break
		}
	}
	fmt.Println(rs)
	//( 0x28
	//) 0x29
}
