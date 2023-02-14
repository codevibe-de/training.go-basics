package main

import (
	"fmt"
)

func simpleChannel() {
	ch := make(chan int)
	// start a goroutine
	ch <- 1
}

func main() {
	ch := make(chan string)
	go sayHello(ch)
	ch <- "Tom"
	ch <- "Jules"
}

func sayHello(ch chan string) {
	for {
		name := <-ch
		fmt.Printf("Hello %s\n", name)
	}
}
