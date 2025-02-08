package main

import "fmt"

func switch1() {
	n := 2
	switch n {
	case 0, 1:
		fmt.Println("It's a zero or a one")
	case 2:
		fmt.Println("It's a two")
	default:
		fmt.Println("Some number other than zero, one or two")
	}
}

func switch2() {
	n := 2
	switch {
	case n == 1:
		fmt.Println("It's a one")
	case n >= 2:
		fmt.Println("It's two or more")
	default:
		fmt.Println("It's a number less than 1")
	}
}

func switch3() {
	n := 43
	switch mod := n % 2; mod {
	case 0:
		fmt.Printf("%d is an even number\n", n)
	default:
		fmt.Printf("%d is NOT an even number\n", n)
	}
}
