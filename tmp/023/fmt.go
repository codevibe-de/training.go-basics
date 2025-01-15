package main

import "fmt"

func main() {
	fmt.Println("Formatting:")
	var n = 42
	fmt.Printf("%d %04d %b %x %T", n, n, n, n, n)
}
