package main

import (
	"fmt"
	"strings"
)

func validator(word string) bool {

	word = strings.ToLower(word)

	chars := make(map[rune]bool)

	for _, val := range word {

		if _, ok := chars[val]; ok {
			return false
		}
		chars[val] = true
	}

	return true
}

func main() {
	words := []string{
		"Hello World",
		"Insomnia",
		"Batman",
		"cat",
	}

	for _, word := range words {
		result := validator(word)
		fmt.Printf("%q -> %v\n", word, result)
	}
}
