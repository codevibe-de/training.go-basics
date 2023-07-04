package main

import "sync"

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (sc *SafeCounter) inc() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count++
}

func main() {
	sc := new(SafeCounter)
	go func() { sc.inc() }()
	go func() { sc.inc() }()
	go func() { sc.inc() }()
}
