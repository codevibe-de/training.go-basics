package main

import (
	"fmt"
)

func main() {
	m := make(map[string]bool)
	m["hey"] = true
	m["ho"] = false
	fmt.Println(m, len(m))
	delete(m, "ho")
	fmt.Println(m, len(m))

	v := m["blah"]
	_, e := m["blah"]
}
