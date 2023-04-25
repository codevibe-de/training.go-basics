package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"time"
)

func consumer(id int, ch chan int) {
	for n := range ch {
		fmt.Printf("#%d got %d\n", id, n)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("Consumer %d done\n", id)
}

func eitherOr(chint chan int, chstr chan string) {
	for {
		select {
		case i := <-chint:
			fmt.Println("got int", i)
		case s := <-chstr:
			fmt.Println("got string", s)
		case chstr <- "hello":
			fmt.Println("saying hello")
		}
	}
}

func miner(quitCh chan struct{}) {
	const hashesPerLoop = 1000000
	var hashCount int64 = 0
	startMillis := time.Now().UnixMilli()
	for {
		select {
		case <-quitCh:
			fmt.Println("Time to quit")
			return
		default:
			for i := 0; i < hashesPerLoop; i++ {
				hasher := md5.New()
				s := fmt.Sprintf("blahbvlahblah %d", i)
				strReader := strings.NewReader(s)
				io.Copy(hasher, strReader)
				hashCount++
				// fmt.Printf("%x\n", hasher.Sum(nil))
			}
			expiredMillis := time.Now().UnixMilli() - startMillis
			hashRate := hashCount / expiredMillis
			fmt.Printf("Current hashrate: %d H/ms\n", hashRate)
		}
	}
}

func main() {
	quitCh := make(chan struct{})
	go miner(quitCh)
	time.Sleep(5 * time.Second)
	quitCh <- struct{}{}
	time.Sleep(time.Second)
}
