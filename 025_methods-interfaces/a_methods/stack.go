package main

import "fmt"

type Stack struct {
	items []string
}

func stackDemo() {
	s := NewStack()
	s.Push("first")
	s.Push("second")
	fmt.Println(s.Peek()) // "second"
	s.Push("third")
	fmt.Println(s.Pop())  // "third"
	fmt.Println(s.Pop())  // "second"
	fmt.Println(s.Peek()) // "first"
}

func NewStack() *Stack {
	s := new(Stack)
	s.items = make([]string, 0, 10)
	return s
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	l := len(s.items)
	if l == 0 {
		return ""
	} else {
		item := s.Peek()
		s.items = s.items[:l-1]
		return item
	}
}

func (s *Stack) Peek() string {
	return s.items[len(s.items)-1]
}
