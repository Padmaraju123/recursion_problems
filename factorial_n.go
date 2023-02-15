package main

import "fmt"

func factorial(n int)int{
	if n==1{
		return 1
	}
	return n*factorial(n-1)


}

func main(){
	var n int
	fmt.Scanln(&n)
	fmt.Println(factorial(n))
}