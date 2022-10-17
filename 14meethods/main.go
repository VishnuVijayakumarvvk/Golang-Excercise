package main

import "fmt"

func main() {
	fmt.Println("This program is based on struct-----")
	Vishnu := User{"Vishnu", 24, "M", true}
	fmt.Println("The struct is : ", Vishnu)
	fmt.Printf("The struct details are : %+v\n", Vishnu)
	fmt.Printf("The name is %v and the age is %v and gender is %v\n", Vishnu.Name, Vishnu.Age, Vishnu.Gender)
	fmt.Println("----------------------")
	Vishnu.Adder()
	fmt.Printf("The name is %v and the age is %v and gender is %v\n", Vishnu.Name, Vishnu.Age, Vishnu.Gender)

}

type User struct {
	Name      string
	Age       int
	Gender    string
	Available bool
}

func (u User) Adder() {
	u.Name = "Vijayakumar"
	fmt.Println("The username is: ", u.Name)
}
