package main

import (
	"fmt"
	"math"
)

func forDemo() {
	var squares [5]int
	for i := 0; i < len(squares); i++ {
		squares[i] = i * i
	}
	fmt.Println(squares)
}

func forDemo2() {
	x := 100.0
	for x > 2 {
		x = math.Sqrt(x)
		fmt.Printf("%.3f ", x)
	}
}

func forDemoRange() {
	números := [4]string{"cero", "uno", "dos", "tres"}
	for i, v := range números {
		fmt.Printf("#%d %s, ", i, v)
	}
}
