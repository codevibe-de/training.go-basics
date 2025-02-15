package main

import "fmt"

type operatorFunc func(int, int) int

func main() {
	fmt.Println("1st class Functions:")

	var sumOperator operatorFunc = func(a, b int) int { return a + b }
	var minusOperator operatorFunc = func(a, b int) int { return a - b }
	var multiplyOperator operatorFunc = func(a, b int) int { return a * b }

	oneMinusOne := minusOperator(1, 1)
	fmt.Println(oneMinusOne)

	twoPlusFourTimesTen := multiplyOperator(sumOperator(2, 4), 10)
	fmt.Println(twoPlusFourTimesTen)
}
