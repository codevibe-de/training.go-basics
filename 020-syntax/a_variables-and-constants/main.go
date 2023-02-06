package main

import "fmt"

func main() {
	fmt.Println("Variables & constants:")
	constants()
	iotaConsts()
}

func constants() {
	const x = 254
	const y int = 1
	const (
		a = "Hello"
		b = "World"
	)
	fmt.Println(x, y, a, b)
	fmt.Printf("Constant `x` has type ´%T`\n", x)
	var d8 uint8 = x
	var d16 int16 = x
	var d32 int32 = x
	var d64 int64 = x
	fmt.Println(d8, d16, d32, d64)

	var v8 int8 = 1
	//var v16 int16 = v8 // doesn't compile
	fmt.Println(v8)
}

func iotaConsts() {
	const (
		_ = (1 << (10 * iota))
		KB
		MB
		GB
		TB
	)
	fmt.Println(KB, MB, GB)
	fmt.Printf("1 TB has %d bytes\n", TB)
}
