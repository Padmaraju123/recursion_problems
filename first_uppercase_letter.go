package main

import (
	"fmt"
	"unicode"
)

func findout(word string) rune {
	var out rune
	for _, v := range word {
		if unicode.IsUpper(v) {
			out = v
			break
		}
	}
	return out
}

func main() {
	word := "wesdHsds"
	gh := findout(word)
	fmt.Printf("%T %c\n", gh, gh)
}
