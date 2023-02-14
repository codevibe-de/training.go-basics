package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 1
		ch <- 2
	}(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
