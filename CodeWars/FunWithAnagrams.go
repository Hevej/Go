package main

import (
	"fmt"
	"sort"
	"strings"
)

func isEqual(a string, b string) bool {
	if len(a) == len(b) {
		sortA := strings.Split(a, "")
		sortB := strings.Split(b, "")
		sort.Strings(sortA)
		sort.Strings(sortB)
		for i, letterA := range sortA {
			if letterA == sortB[i] {
				continue
			} else {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func isIn(text []string, a string) bool {
	for _, word := range text {
		if word == a {
			return true
		}
	}
	return false
}

func funWithA(text []string) []string {
	sort.Strings(text)
	keepWord := text[0]
	var aux []string
	aux = append(aux, keepWord)
	for i := 1; i < len(text); i++ {
		if isEqual(keepWord, text[i]) {
			continue
		} else {
			keepWord = text[i]
			aux = append(aux, text[i])
		}
	}
	return aux
}

func main() {
	text := []string{"code", "bbb", "odec"}
	result := funWithA(text)
	fmt.Println(result)
}
