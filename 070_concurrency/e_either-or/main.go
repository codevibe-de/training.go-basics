package main

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int)
	go func(ch chan int) {
		ticker := time.NewTicker(800 * time.Millisecond)
		var n = 0
		for range ticker.C {
			n++
			ch <- n
		}
	}(intCh)

	stringCh := make(chan string)
	go func(ch chan string) {
		ticker := time.NewTicker(1100 * time.Millisecond)
		for range ticker.C {
			ch <- time.Now().String()
		}
	}(stringCh)

	for {
		select {
		case n := <-intCh:
			fmt.Println("Got an int: ", n)
		case s := <-stringCh:
			fmt.Println("Got a string: ", s)
		}
	}
}
