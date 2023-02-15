package main

import (
	"fmt"
	"sync"
	"time"
)

func foo() {
	fmt.Println("Hello")
	time.Sleep(2 * time.Second)
}

func newDemo() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		fmt.Println("Go go go")
		wg.Done()
	}()
	go func() {
		fmt.Println("Go go go!!")
		wg.Done()
	}()
	wg.Wait()
}

func main() {
	wg := new(sync.WaitGroup)
	channel := make(chan string) // capacity 0
	wg.Add(1)
	go sayHello(channel, wg)

	channel <- "Robert"
	channel <- "Vanessa"
	channel <- "Markus"
	close(channel)
	wg.Wait()
}

func sayHello(ch chan string, wg *sync.WaitGroup) {
	for name := range ch {
		fmt.Println("Hello", name)
	}
	wg.Done()
}
