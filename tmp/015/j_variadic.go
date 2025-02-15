package main

import "fmt"

func printNames(intro string, names ...string) {
	fmt.Println(intro)
	fmt.Println(names)
}

func main() {
	printNames("Unsere Helden sind: ", "Tim", "Struppi")
}

// Unsere Helden sind:
// [Tim Struppi]
