package main

import "fmt"

func main() {
	s := "托"
	fmt.Println(len(s))

	bytes := []byte(s)
	for i, v := range bytes {
		fmt.Println(i, v)
	}
}

// 3
// 0 230
// 1 137
// 2 152
