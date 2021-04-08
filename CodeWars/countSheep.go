package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	num, _ := strconv.Atoi(args[1])
	var sheeps string = ""
	for i := 1; i <= num; i++ {
		sheeps = strings.Join([]string{sheeps, strconv.Itoa(i)}, "")
		sheeps = strings.Join([]string{sheeps, "sheep..."}, " ")
	}
	fmt.Println(sheeps)
}
