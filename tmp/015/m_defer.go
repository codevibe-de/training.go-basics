package main

import "fmt"

func deferDemo() {
	defer fmt.Println("you're gonna be fine")
	fmt.Println("Sooner or")
	defer func() {
		fmt.Println("later")
	}()
}
