package main

import "fmt"

func find_cons(word string)int{
	if word[0]!=97 || word[0]!=101 || word[0]!=105 || word[0]!=111 || word[0]!=117{
		return 0
	}else{
		return 1
	}
	return find_cons(word[1:])
}


func main() {
	word := "padamraju"
	fmt.Println(find_cons(word))
}
