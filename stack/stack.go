package stack

import (
	"sync"
)

type Stack struct {
	sync.Mutex
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

func (s *Stack) Push(v int) {
	s.Lock()
	defer s.Unlock()

	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, bool) {
	s.Lock()
	defer s.Unlock()

	if len(s.items) > 0 {
		v := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return v, true
	}

	// stack is empty
	return 0, false
}
