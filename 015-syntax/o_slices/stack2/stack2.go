package stack2

import "fmt"

type Stack []string

func Demo() {
	fmt.Println("Slices:")

	myStack := make(Stack, 0)
	fmt.Println(myStack)

	Push(&myStack, "first")
	fmt.Println(myStack)

	Push(&myStack, "second")
	item, _ := Pop(&myStack)
	fmt.Printf("popped: '%s'\n", item)
	fmt.Println(myStack)

	Push(&myStack, "third")
	fmt.Println(myStack)
	fmt.Println("peeked:", Peek(&myStack))
}

func Push(st *Stack, item string) {
	*st = append(*st, item)
}

func Pop(st *Stack) (string, bool) {
	l := len(*st)
	if l == 0 {
		return "", false
	} else {
		popped := (*st)[l-1]
		*st = (*st)[:l-1]
		return popped, true
	}
}

func Peek(st *Stack) string {
	return (*st)[len(*st)-1]
}
