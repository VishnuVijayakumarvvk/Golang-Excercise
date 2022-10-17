package main

import "fmt"

func main() {
	defer fmt.Println("one")
	fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("four")
	defer fmt.Println("five")
	deferstd()

}
// two four five three one

func deferstd()  {
	for i := 0; i <=5 ; i++ {
		 defer fmt.Println(i)
	}
} 
//012345
// two four 543210 five three one