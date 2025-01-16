package main

import (
	"fmt"
	"time"
)

// magic date: 02. Januar 2006, 15:04:05

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Printf("%02d.%02d.%04d\n", now.Day(), now.Month(), now.Year())
	fmt.Println(now.Format("02.01.2006"))

	tz, _ := time.LoadLocation("Australia/Sydney")
	sydneyTime := now.In(tz)
	fmt.Println(sydneyTime)

	fiveMinEarlier := sydneyTime.Add(time.Minute * -5)
	fmt.Println(fiveMinEarlier)

	fmt.Println(fiveMinEarlier.Before(sydneyTime))
}
