package main

import "fmt"

func main() {
	// myRollNumber := 45
	// var ptr = &myRollNumber
	// fmt.Println("the address inn ptr ", ptr)
	// fmt.Println("The value in ptr: ", *ptr)
	// fmt.Println("The sum: ", *ptr + 5)

	number := 50
	var ptr = &number
	fmt.Println("The address of number: ", ptr)
	fmt.Println("The value : ", *ptr)
	fmt.Println("The sum is: ", *ptr*2)
}
