package main

import "fmt"

func find_len(word string)int{
	if word == ""{
		return 0
	}
	return 1+find_len(word[1:])

}

func main(){
	word := "padmaraju"
	fmt.Println(find_len(word))
}