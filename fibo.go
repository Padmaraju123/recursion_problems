
//print the n fibonacci numbers
package main

import "fmt"

func fibbbooo(n int){
	a:=0
	b:=1

	if n==1{
		fmt.Println(a)
		return
	}
	if n==2{
		fmt.Println(a)
		fmt.Println(b)
		return
	}

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
	fmt.Println("enter the n value...")
	var n int
	fmt.Scanln(&n)

	fibbbooo(n)
}
