package main

import "fmt"

func main() {
	fmt.Println("If-Statements:")
	numbers := [10]int{2, 22, 333, 4, -123, 888, 4, 8080}
	min, max := findMinMax(numbers)
	fmt.Printf("min=%d, max=%d", min, max)
}

func findMinMax(arr [10]int) (int, int) {
	min, max := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, max
}
