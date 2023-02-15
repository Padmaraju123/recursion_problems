package main

import "fmt"

func factorial(num int)int{
	if num == 0 || num ==1{
		return 1
	}else{
		return num*factorial(num-1)
	}
}

func main(){
	var num int
	fmt.Println("enter the number to factorial....")
	fmt.Scanln(&num)
	fmt.Println(factorial(num))
}