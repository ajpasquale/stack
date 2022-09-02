package stack

type stack struct {
	items []int
}

func newStack(size int) *stack {
	return &stack{
		items: make([]int, 0, size),
	}
}

func (s *stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *stack) Pop() {
	s.items = s.items[:len(s.items)-1]

}

func (s *stack) Peek() int {
	return s.items[len(s.items)-1]
}

func (s *stack) Size() int {
	return len(s.items)
}
