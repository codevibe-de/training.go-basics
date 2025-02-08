package main

import "fmt"

func pointer() {
	var strPointer *string
	text := "abc"
	strPointer = &text
	fmt.Println(strPointer, *strPointer)
}

func pointer2() {
	text := "abc"
	modify(&text)
	fmt.Println(text)
}

func modify(s *string) {
	*s = "new value"
}
