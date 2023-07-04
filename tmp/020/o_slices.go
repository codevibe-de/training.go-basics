package main

import "fmt"

func main() {
	sliceLenAndCap()
}

func sliceDemo() {
	intSlice := []int{1, 2, 3}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)
}

func sliceLenAndCap() {
	intSlice := []int{1, 2, 3}
	intSlice = append(intSlice, 4)
	fmt.Printf("length: %d, capacity: %d", len(intSlice), cap(intSlice))
	// length: 4, capacity: 6
}

func slicing() {
	boolArray := [4]bool{true, true, false, false}
	boolSlice := boolArray[1:3]
	fmt.Println(boolSlice)
}

func makeSlice() {
	var strSlice []string = make([]string, 2, 10)
	fmt.Printf("%q\n", strSlice)
	fmt.Printf("length: %d, capacity: %d", len(strSlice), cap(strSlice))
}
