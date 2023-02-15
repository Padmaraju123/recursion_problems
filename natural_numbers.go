//given n print the values upto 1 
package main

import "fmt"

func printing(n int){
	if n<1{
		return 
	}
	printing(n-1)
	fmt.Println(n)

}

func main(){
	var n int
	fmt.Scanln(&n)
	printing(n)

}