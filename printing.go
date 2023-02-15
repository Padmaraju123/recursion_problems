package main

import "fmt"

func print5(n int)int{
	return n+1
}


func print4(n int)int{
	return print5(n+1)+1
}


func print3(n int)int{
	return print4(n+1)+1
	
}


func print2(n int)int{
	return print3(n+1)+1
}


func print1(n int)int{
	return print2(n+1)

}

func main(){
	fmt.Println(print1(1))
}
