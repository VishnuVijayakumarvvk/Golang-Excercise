package main

import "fmt"

func main() {
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "orange"
	fruitlist[3] = "banana"
	fmt.Println("The fruit array is : ", fruitlist)
	fmt.Println("The length of fruitlist is: ", len(fruitlist))

	var veglist = [4]string{"potato", "cabbage", "beans"}
	fmt.Println("The veg list is: ", veglist)
	fmt.Println("The length of fruitlist is: ", len(veglist))

}
