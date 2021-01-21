package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide one text file to process!")
		os.Exit(1)
	}

	filename := arguments[1]
	open, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s", err)
		os.Exit(1)
	}
	defer open.Close()

	//Holds number of linesin the input file that did not match any one of two regular expressions
	notAMatch := 0
	r := bufio.NewReader(open)
	for {
		line, err := r.ReadString('n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
		}

		//Parse a regular expression and returnm if successful, a Regexp object
		//Panics if the expression cannot be parsed
		r1 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)

		//MatchString dice si en esa string hay alguna expresion regular que haga match
		if r1.MatchString(line) {
			//FindStringSubmatch returns a slice of strings
			//La exp regular se encuentra en la lista devuelta
			match := r1.FindStringSubmatch(line)
			for i, ma := range match {
				fmt.Println(i, "-", ma)
			}
			d1, err := time.Parse("02/Jan/2006:15:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}

		r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`)

		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}
	}
	fmt.Println(notAMatch, "lines did not match!")
}
