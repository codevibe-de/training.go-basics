package main

import "fmt"

func main() {
	slc1 := []string{"one", "twó", "three"}
	slc2 := []string{"one", "two", "three"}

	same := compareSlices(slc1, slc2)
	fmt.Printf("Are the slices the same? %t", same)
}

func compareSlices(s1, s2 []string) bool {
	// TODO this is not optimal!
	return len(s1) == len(s2)
}
