package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go sayHello(ch)
	ch <- "Tom"
	ch <- "Jules"
}

func sayHello(ch chan string) {
	name := <-ch // lesen ist auch BLOCKIEREND
	fmt.Printf("Hello %s\n", name)
}
