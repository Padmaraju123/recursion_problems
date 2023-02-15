//given a number write a program to sum of the digits of the number

package main

import "fmt"

func sum_of_digits(num int)int{
	if num<10{
		return num
	}
	vv := num%10
	num = num/10
	return vv+sum_of_digits(num)
}

func main(){
	var num int
	fmt.Scanln(&num)
	fmt.Println(sum_of_digits(num))
}