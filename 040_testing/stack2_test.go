package main

import (
	"fmt"
	"local/020-syntax/o_slices/stack2"
)

func ExamplePush() {
	st := make(stack2.Stack, 0)
	stack2.Push(&st, "foo")
	fmt.Printf("%q\n", st)
	// Output: ["foo"]
}
