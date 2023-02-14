package main

import (
	"fmt"
	"time"
)

type workdata struct {
	a      int
	b      int
	result int
}

func add(a, b int) int {
	time.Sleep(time.Second)
	return a + b
}

func work(inCh chan workdata, outCh chan workdata) {
	for {
		wd := <-inCh
		wd.result = add(wd.a, wd.b)
		outCh <- wd
	}
}

func main() {
	inCh := make(chan workdata)
	outCh := make(chan workdata)
	go work(inCh, outCh)
	inCh <- workdata{a: 10, b: 20}
	inCh <- workdata{a: -2, b: 4}
	fmt.Println(<-outCh)
	fmt.Println(<-outCh)
}
