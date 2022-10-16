package main

import "fmt"

func main() {
	var number = 4
	if number < 3 {
		fmt.Println("The number is less that 3")
	} else if number == 4 {
		fmt.Println("The number is equal to 4")
	} else {
		fmt.Println("The number is grater than 4")
	}

	fmt.Println("End of if statement---------")

	if no := 3; no < 4 {
		fmt.Println("The number is less than 4")
	}
	fmt.Println("===============")

	if kg := 4; number < 6 {
		fmt.Println("The number is grater than kg", kg)
	}
}
