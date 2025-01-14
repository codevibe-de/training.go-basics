package main

import "fmt"

func funcVars() {
	// define function type
	type filterFunc func(string) bool
	// implement it
	var isEmpty filterFunc = func(s string) bool {
		return len(s) == 0
	}
	// call it
	fmt.Println(isEmpty(""), isEmpty("abc"))
}

// Output: true false

type filterFunc func(string) bool

func main() {
	atLeastThree := createFilterFunc(3)
	fmt.Println(atLeastThree("ab"), atLeastThree("abc"))
}

func createFilterFunc(length int) filterFunc {
	return func(s string) bool {
		return len(s) >= length
	}
}
