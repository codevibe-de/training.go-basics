package main

import (
	"fmt"
	"strings"
)

func stringsDemo() {
	s := "Hey you"
	strings.Split(s, " ")
}

func arrayDemo() {
	var arr [3]int
	fmt.Println(arr)
	arr[1] = 8888
	fmt.Println(arr)

	strArr := [3]string{"Hey you", "Out there in the cold"}
	fmt.Printf("%q\n", strArr)
}

func main() {
	arrayDemo()
}