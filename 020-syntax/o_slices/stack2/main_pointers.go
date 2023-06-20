package main

import "fmt"

type stack []string

func main() {
	fmt.Println("Slices:")

	myStack := make(stack, 0)
	fmt.Println(myStack)

	push(&myStack, "first")
	fmt.Println(myStack)

	push(&myStack, "second")
	item := pop(&myStack)
	fmt.Printf("popped: '%s'\n", item)
	fmt.Println(myStack)

	push(&myStack, "third")
	fmt.Println(myStack)
	fmt.Println("peeked:", peek(&myStack))
}

func push(st *stack, item string) {
	*st = append(*st, item)
}

func pop(st *stack) string {
	l := len(*st)
	if l == 0 {
		return ""
	} else {
		popped := (*st)[l-1]
		*st = (*st)[:l-1]
		return popped
	}
}

func peek(st *stack) string {
	return (*st)[len(*st)-1]
}
