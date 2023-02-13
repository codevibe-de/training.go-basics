package main

import (
	"fmt"
	"testing"
)

func TestNextFib(t *testing.T) {
	Reset()
	wanted := []int{1, 1, 2, 3, 5, 8, 13}
	for n := 0; n < len(wanted); n++ {
		fib := NextFib()
		if fib != wanted[n] {
			t.Errorf("Fibonacci no. %d: Got %v, wanted %v", n, fib, wanted[n])
		}
	}
}

func ExampleNextFib() {
	Reset()
	fmt.Println(NextFib())
	// Output: 1
}
