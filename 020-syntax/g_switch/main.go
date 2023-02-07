package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch:")
	checkForWeekend()
}

func checkForWeekend() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
