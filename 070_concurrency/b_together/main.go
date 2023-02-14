package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func startUp(id int) {
	fmt.Printf("#%d starting up ...\n", id)
	seconds := rand.Intn(10) + 1
	fmt.Println(seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("#%d ready!\n", id)
}

func runner(id int, ch chan bool, wg *sync.WaitGroup) {
	startUp(id)
	wg.Done()
	<-ch
	fmt.Println(time.Now())
}

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	signalCh := make(chan bool)
	wg := new(sync.WaitGroup)
	for n := 0; n < 3; n++ {
		wg.Add(1)
		go runner(n, signalCh, wg)
	}
	wg.Wait()
	for n := 0; n < 3; n++ {
		signalCh <- true
	}

	time.Sleep(time.Second)
}
