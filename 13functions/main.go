package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of functons----")
	greeter()
	result := adder(3, 5)
	fmt.Println("The sum is: ", result)

	proresult, err := proadder(2, 4, 6, 3)
	fmt.Println("The proadder result is: ", proresult)
	fmt.Println("The proadder result is: ", err)
}

func greeter() {
	fmt.Println("Hi namastey to the function class")
}

func adder(valo int, valt int) int {
	return valo + valt
}

func proadder(values ...int) (int, string) {
	total := 0
	for i := range values {
		total += values[i]
	}
	return total, "Error"
}
