package main

import (
	"fmt"
	"math/rand"
	"time"
)

func startUp() {
	fmt.Println("Starting up ...")
	seconds := rand.Intn(5) + 1
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Ready!")
}
