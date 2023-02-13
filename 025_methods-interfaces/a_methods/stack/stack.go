package stack

type Stack struct {
	items []string
}

func NewStack() *Stack {
	s := new(Stack)
	s.items = make([]string, 0, 10)
	return s
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top-most item on the stack.
// The boolean return value shows if a value was present, similiar to how a map works.
func (s *Stack) Pop() (item string, found bool) {
	item, found = s.Peek()
	if found {
		s.items = s.items[:len(s.items)-1]
	}
	return item, found
}

func (s *Stack) Peek() (string, bool) {
	l := len(s.items)
	if l == 0 {
		return "", false
	} else {
		return s.items[l-1], true
	}
}
