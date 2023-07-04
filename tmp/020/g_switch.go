package main

import "fmt"

func switchDemo() {
	n := 43
	switch mod := n % 2; mod {
	case 0:
		fmt.Printf("%d is an even number\n", n)
	default:
		fmt.Printf("%d is NOT an even number\n", n)
	}
}
