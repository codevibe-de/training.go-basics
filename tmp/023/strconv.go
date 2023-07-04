package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("123")
	fmt.Println(n, err)

	fmt.Println(strconv.Itoa(42))

	b, err := strconv.ParseBool("true")
	fmt.Println(b, err)

	s := strconv.Quote("foo")
	fmt.Println(s)
}

// 123 <nil>
// 42
// true <nil>
// "foo"
