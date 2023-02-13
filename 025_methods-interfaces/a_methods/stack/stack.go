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
func (s *Stack) Pop() (string, bool) {
	l := len(s.items)
	if l == 0 {
		return "", false
	} else {
		item := s.Peek()
		s.items = s.items[:l-1]
		return item, true
	}
}

func (s *Stack) Peek() string {
	return s.items[len(s.items)-1]
}
