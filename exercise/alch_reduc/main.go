package main

import (
	"fmt"
	"io"
	"local/025_methods-interfaces/a_methods/stack"
	"os"
	"time"
)

func main() {
	var startTime = time.Now()
	if file, err := os.Open("alchemical-data.txt"); err == nil {
		var st = stack.NewStack()
		var b = make([]byte, 1)

		for _, err := file.Read(b); err != io.EOF; _, err = file.Read(b) {
			s := string(b)
			peek, exists := st.Peek()
			if exists && canReact(s, peek) {
				st.Pop()
			} else {
				st.Push(s)
			}
		}
		fmt.Println("Result size", st.Len())
		fmt.Println(time.Now().Sub(startTime))
	} else {
		fmt.Println(err)
	}
}

func canReact(s1, s2 string) bool {
	b1 := ([]byte(s1))[0]
	b2 := ([]byte(s2))[0]
	diff := int(b1) - int(b2)
	return diff == 32 || diff == -32
}
