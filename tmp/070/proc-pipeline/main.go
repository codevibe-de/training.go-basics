package main

import (
	"fmt"
	"time"
)

func main() {
	doneCh := make(chan struct{})

	genOutCh := generator(doneCh)
	filterOutCh := filter(genOutCh, func(n int) bool { return n%2 == 0 })
	mapOutCh := mapper(filterOutCh)
	result := collector(mapOutCh)

	time.Sleep(time.Second)
	doneCh <- struct{}{}
	fmt.Println(result)
}

func gen(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("Generator stage starting")
		i := 0
		for {
			select {
			case ch <- i:
				time.Sleep(100 * time.Millisecond)
				i++
			case <-done:
				fmt.Println("Generator stage stopping")
				close(ch)
				return
			}
		}
	}()
	return ch
}

func filter(inCh chan int, filterFunc func(int) bool) chan int {
	outCh := make(chan int)
	go func() {
		fmt.Println("Filter stage starting")
		for n := range inCh {
			if filterFunc(n) {
				outCh <- n
			}
		}
		fmt.Println("Filter stage stopping")
		close(outCh)
	}()
	return outCh
}

func mapper(inCh chan int) chan int {
	outCh := make(chan int)
	go func() {
		fmt.Println("Mapper stage starting")
		for n := range inCh {
			outCh <- n * n
		}
		fmt.Println("Mapper stage stopping")
		close(outCh)
	}()
	return outCh
}

func collector(inCh chan int) []int {
	res := make([]int, 0, 10)
	go func() {
		fmt.Println("Collector stage starting")
		for n := range inCh {
			res = append(res, n)
		}
		fmt.Println("Collector stage stopping")
	}()
	return res
}
