package main

import "fmt"

func main() {
	fmt.Println("Formatting:\n")
	var n = 42
	fmt.Printf("%d %04d %b %x %T", n, n, n, n, n)
}
