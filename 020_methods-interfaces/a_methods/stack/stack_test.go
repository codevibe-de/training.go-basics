package stack

import (
	"fmt"
	"testing"
)

func TestPushAndPeek(t *testing.T) {
	// given
	stack := NewStack()
	want := "öäü"
	// when
	stack.Push(want)
	gotItem, gotExists := stack.Peek()
	// then
	if gotItem != want || gotExists == false {
		t.Errorf("Push(%s), then Peek(): got (%q,%t), want (%q,%t)", want, gotItem, gotExists,
			want, true)
	}
}

func ExampleStack_Pop() {
	// given
	stack := NewStack()
	stack.Push("a")
	stack.Push("b")
	// when
	item, found := stack.Pop()
	// then
	fmt.Println(item, found)
	// Output: b true
}
