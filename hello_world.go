//print anything with desired number

package main

import "fmt"

func message(i int){
	if i==5{
		fmt.Println(i)
		return 
	}
	fmt.Println(i)
	message(i+1)

}

func main(){

	var i int
	fmt.Println("enter the number....")
	fmt.Scanln(&i)
	message(i)
}