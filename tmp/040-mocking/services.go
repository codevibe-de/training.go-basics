package mock_demo

import "fmt"

func Sum(a, b int, c Calculator) int {
	fmt.Printf("Now calculating the sum of %d and %d", a, b)
	return c.Sum(a, b)
}
