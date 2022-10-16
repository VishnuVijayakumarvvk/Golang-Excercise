package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops----")
	//days := []string{"Sunday","monday","tuesday","wednesday","thursday","friday"}
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println("The day is :",days[i])
	// }

	// for i := range days{
	// 	fmt.Println("The day is ",days[i])
	// }

	// for i,day :=range days{
	// 	fmt.Printf("The index is %v and the day is %v\n",i,day)
	// }

	var number = 1

	for number <= 10 {
		if number == 5 {
			fmt.Println("The number is skipped")
			number++
			continue
		}
		fmt.Println("The number is : ", number)
		number++
	}
}
