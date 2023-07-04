package main

import "fmt"

func withSameTypeParams(a, b, c int) {
}

func returningMultipleValues(name string) (int, error) {
	return 0, fmt.Errorf("function not implemented")
}

func namedReturn(a, b, c, d int) (min, max int) {
	min, max = a, b
	return min, max
}
