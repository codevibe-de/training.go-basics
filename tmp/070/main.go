package main

import (
	"fmt"
	"time"
)

func simpleChannel() {
	ch := make(chan int)
	// start a goroutine
	ch <- 1
}

func bufferedChannel() {
	ch := make(chan int, 100)
	fmt.Printf("Channel contains %d of max %d items\n", len(ch), cap(ch))
	// we can write into channel although there
	// is no consuming goroutine
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Printf("Channel contains %d of max %d items\n", len(ch), cap(ch))
}

func closeChannel() {
	dummyChannel := make(chan int)
	go func(ch chan int) {
		fmt.Println("Waiting for input...")
		n, exists := <-ch
		fmt.Println(n, exists)
	}(dummyChannel)
	time.Sleep(time.Second)
	close(dummyChannel)
	time.Sleep(time.Second)
}

func readUntilClosed(ch chan int) {
	for {
		input, exists := <-ch
		if !exists {
			fmt.Println("Feierabend!")
			return
		} else {
			fmt.Println(input) // work with input
		}
	}
}

func readUntilClosedRanged(ch chan int) {
	for input := range ch {
		fmt.Println(input) // work with input
	}
}

func main() {
	bufferedChannel()
}
