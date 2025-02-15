package main

import "fmt"

func strSlice() {
	var s = "Hello"
	fmt.Println(s[3:])
}

func stringRange() {
	s := "números"
	fmt.Println("idx | unicode | char")
	fmt.Println("----+---------+-----")
	for i, c := range s {
		fmt.Printf(" #%d |     %d | %q\n", i, c, c)
	}
}

func isCaptialA(r rune) bool {
	return r == 65
}

func runes() {
	s := "Thomas 托马斯"
	runes := []rune(s)
	fmt.Println(runes)

	s2 := string([]rune{'托', 'H', 'i'})
	fmt.Println(s2)

	for i, v := range s2 {
		fmt.Printf("#%d %v\n", i, v)
	}
}
