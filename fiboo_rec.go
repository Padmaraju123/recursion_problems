package main

import "fmt"

func fiboo_rec(n int)int{
	if n == 1 || n==2{
		return 1
	}
	return fiboo_rec(n-1)+fiboo_rec(n-2)
}

func main(){
	fmt.Println("enter the n value...")
	var n int
	fmt.Scanln(&n)
	fmt.Println(fiboo_rec(n))
}