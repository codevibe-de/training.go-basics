package main

import "fmt"

func iotaDemo() {
	const (
		zero = iota * 2
		two  = iota * 2
		four
		_
		eight
	)
	fmt.Println(zero, two, four, eight)
}
