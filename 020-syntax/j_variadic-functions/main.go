package main

import "fmt"

func main() {
	fmt.Println("Variadic Functions:")
	a, b, c, d := 10, 2, 0, -42
	doubleInts(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
}

func doubleInts(ints ...*int) {
	for i := range ints {
		*ints[i] = *ints[i] * 2
	}
}
