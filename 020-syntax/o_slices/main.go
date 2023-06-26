package main

import (
	"fmt"
)

type stack []string

func main() {
	demoStack1()
	//stack2.Demo()
}

func demoStack1() {
	fmt.Println("Slices:")

	myStack := make(stack, 0)
	fmt.Println(myStack)

	myStack = push(myStack, "first")
	fmt.Println(myStack)

	myStack = push(myStack, "second")
	myStack, item := pop(myStack)
	fmt.Printf("popped: '%s'\n", item)
	fmt.Println(myStack)

	myStack = push(myStack, "third")
	fmt.Println(myStack)
}

func push(st stack, item string) stack {
	return append(st, item)
}

func pop(st stack) (stack, string) {
	l := len(st)
	if l == 0 {
		return st, ""
	} else {
		return st[:l-1], st[l-1]
	}
}

func peek(st stack) string {
	return st[len(st)-1]
}
