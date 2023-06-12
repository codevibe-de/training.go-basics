package main

import "fmt"

// https://www.frontlinesedutech.com/blog/why-computers-screw-up-doing-math-on-floating-point-numbers
func main() {
	fmt.Println("For-Loops:")
	total := sum([]float32{0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1})
	fmt.Println(total)
}

func sum(numbers []float32) float32 {
	var result float32 = 0.0
	for _, v := range numbers {
		result += v
	}
	return result
}
