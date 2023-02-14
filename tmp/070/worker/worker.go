package main

import (
	"fmt"
	"runtime"
	"time"
)

type workdata struct {
	a      int
	b      int
	result int
}

func worker(inCh <-chan workdata, outCh chan<- workdata) {
	for wd := range inCh {
		wd.result = add(wd.a, wd.b) // some tough calculation stuff
		outCh <- wd
	}
}

func reporter(outCh <-chan workdata) {
	for {
		fmt.Println(<-outCh)
	}
}

func add(a, b int) int {
	fmt.Printf("Calculating %d + %d...\n", a, b)
	time.Sleep(100 * time.Millisecond)
	return a + b
}

func main() {
	fmt.Printf("Running 2 workers on system with %d cores\n", runtime.NumCPU())
	inCh := make(chan workdata)
	outCh := make(chan workdata)
	go worker(inCh, outCh)
	go worker(inCh, outCh)
	go reporter(outCh)
	inCh <- workdata{a: 10, b: 20}
	inCh <- workdata{a: -2, b: 4}
	close(inCh) // that's it
	fmt.Println("Work has been handed out - waiting for results to come in")
	time.Sleep(time.Second * 4)
}
