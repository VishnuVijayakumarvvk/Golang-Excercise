package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice:) ")
	var fruitlist = []string{"Apple", "Orange", "banana"}
	fruitlist = append(fruitlist, "beans", "Dragon fruit", "jack Fruit")
	fmt.Println(fruitlist)
	fruitlist = append(fruitlist[3:6])
	fmt.Println(fruitlist)

	rollnumber := make([]int, 3)
	rollnumber[0] = 12
	rollnumber[1] = 4
	//rollnumber[2] = 78
	fmt.Println(rollnumber)
	rollnumber = append(rollnumber, 34, 54, 32)
	fmt.Println(rollnumber)
	sort.Ints(rollnumber)
	fmt.Println(rollnumber)

	fmt.Println("Remove numbers from slice: -----------")
	var rollno = []int{12, 45, 7, 23, 87}
	var index int = 2
	rollno = append(rollno[:index], rollno[index+1:]...)
	fmt.Println("The edited slice is : ", rollno)
}
