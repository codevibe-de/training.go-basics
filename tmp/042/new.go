package main

import "fmt"

func main() {
	var p *int64 = new(int64)
	fmt.Println(*p) // Output: 0

	p2 := new(Messages)
	fmt.Println(*p2) // Output: {[] 0}
}

type Messages struct {
	items             []string
	lastChangedMillis uint64
}
