package main

import "fmt"


func fiboo(n int){
	a:=0
	b:=1
	fmt.Println(a)
	fmt.Println(b)
	for i:=0;i<n-2;i++{
		c:=a+b
		a = b
		b = c
		fmt.Println(c)
	}
}

func main(){
	var n int
	fmt.Scanln(&n)
	fiboo(n)
}