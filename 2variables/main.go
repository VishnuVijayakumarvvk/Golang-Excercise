package main

import "fmt"

const Logginname string="VishnuVij"

func main() {
	var username string = "Vishnu Vijayakumar"
	fmt.Println(username)
	fmt.Printf("The Username type is: %T",username)

	var secondname = "Vishnu"
	fmt.Println(secondname)
	fmt.Printf("The Username type is: %T",secondname) 
	 
	rollnumber := 64
	fmt.Println(rollnumber)
	fmt.Printf("The Username type is: %T",rollnumber)

	fmt.Println("\n")
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)

	var decimalNumber float64=3.146328732892
	fmt.Println(decimalNumber)
	fmt.Println("The type is %T",decimalNumber)

	fmt.Println(Logginname)
	 
	secondName := "Vijayakumar"
	fmt.Println(secondName)
	fmt.Printf("The type is %T",secondName)
 
	Thirdname := "Givdehs"
}
