package main

import "fmt"

func fiboo_rec(n int)int{

	if n<2{
		return n
	}

	return fiboo_rec(n-1)+fiboo_rec(n-2)
}


func main(){
	var n int
	fmt.Scanln(&n)
	fmt.Println(fiboo_rec(n))
}