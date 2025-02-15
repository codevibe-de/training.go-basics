package main

import "fmt"

func main() {
	fmt.Println("Variadic Functions:")
	a, b, c, d := 10, 2, 0, -42
	doubleInts(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
}

func doubleInts(ints ...*int) {
	for _, v := range ints {
		*v *= 2 // oder: *v = *v * 2
	}
}
