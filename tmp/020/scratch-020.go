package main

import "fmt"

func main() {
	strBytes()
	strSlice()
}

func strBytes() {
	s := "Ȫ"
	fmt.Println(len(s))

	fmt.Println(s[1])

	bytes := []byte(s)
	for i, v := range bytes {
		fmt.Println(i, v)
	}

	bytes = []byte{200, 170}
	s = string(bytes)
	fmt.Println(s)
}

// 2
// 170
// 0 200
// 1 170
// Ȫ

func strSlice() {
	var s = "Hello"
	fmt.Println(s[3:])
}

// lo
