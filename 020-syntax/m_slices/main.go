package main

import "fmt"

type stack []string

func main() {
	fmt.Println("Slices:")
	myStack := make(stack, 0)
	myStack = push(myStack, "first")
	myStack = push(myStack, "second")
	myStack = pop(myStack)
	myStack = push(myStack, "third")
	fmt.Println(myStack)
}

func push(st stack, item string) stack {
	return append(st, item)
}

func pop(st stack) (stack, string) {
	item := st[len(st)-1]
	return st[:len(st)-1]
}
