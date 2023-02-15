package main

import (
	"fmt"
	"unicode"
)

func findout_rec(word string, index int) rune {
	if unicode.IsUpper(rune(word[index])) {
		return rune(word[index])
	}
	if index == len(word)-1 {
		fmt.Println("no letter has upper case in given word...")
	}
	index = index + 1
	return findout_rec(word, index)

}

func main() {
	word := "dadfsTRdfd"
	index := 0
	fmt.Printf("%c\n", findout_rec(word, index))
}
