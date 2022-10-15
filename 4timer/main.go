package main

import (
	"fmt"
	"time"
)

func main() {
	timenow := time.Now()
	fmt.Println(timenow)
	timeformat := timenow.Format("01-02-2006 Monday 15:04:05")
	fmt.Println(timeformat)

	examdate := time.Date(2021, time.August, 14, 4, 5, 6, 7,time.UTC)
	fmt.Println(examdate)

	fmt.Println("The time now is: ",timenow.Format("15:04:05"))
}
