package main

import (
	"fmt"
	"time"
)

func main() {
	duration, _ := time.ParseDuration("1h2m3s4ms")
	fmt.Println(duration)

	sinceMillenium := time.Since(
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Now().Location()),
	)
	fmt.Println(sinceMillenium)

	time.AfterFunc(time.Millisecond*100, func() {
		fmt.Println("100 millis have elapsed")
	})
	time.Sleep(time.Second)
}
