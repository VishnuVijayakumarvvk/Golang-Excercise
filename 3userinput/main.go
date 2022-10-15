package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the pizza hotel")
	fmt.Printf("Please enter the pizza rating: ")
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	number,_ := strconv.ParseFloat(strings.TrimSpace(input),64)
	fmt.Println("The added number is : ", number+1)
}
