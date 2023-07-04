package main

import (
	"fmt"
	"time"
)

func main() {
	chOfInts := make(chan int)
	chOfStrings := make(chan string)
	go worker(chOfInts, chOfStrings)
	chOfStrings <- "hello"
	chOfInts <- 1
	time.Sleep(time.Second)
}

func worker(intCh chan int, stringCh chan string) {
	for {
		select {
		case n, ok := <-intCh:
			if ok {
				fmt.Println("Got an int: ", n)
			}
		case s := <-stringCh:
			fmt.Println("Got a string: ", s)
		default:
			fmt.Println("Got nothing")
			// do something useful here
		}
	}
}
