package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("The number is 1")
	case 2:
		fmt.Println("The number is 2")
		fallthrough
	case 3:
		fmt.Println("The number is 3")
	case 4:
		fmt.Println("The number is 4")
	case 5:
		fmt.Println("The number is 5")
	case 6:
		fmt.Println("The number is 6")
		fallthrough
	default:
		fmt.Println("Something went wrong!!...")
	}

	rand.Seed(time.Now().Unix())
	dice := rand.Intn(6) +1
	fmt.Println("The dice number is: ", dice)
}
