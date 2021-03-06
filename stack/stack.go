package stack

import (
	"sync"

	list "github.com/zianwar/datastructures/linkedlist/single"
)

// Stack represents a thread-safe stack data structure
type Stack struct {
	sync.Mutex
	l    *list.List
	size int
}

// NewStack returns new stack
func NewStack() *Stack {
	return &Stack{
		l: list.NewList(),
	}
}

// Size returns the size of the stack
// O(1)
func (s *Stack) Size() int {
	return s.size
}

// Peek returns the item at the top of the stack
// O(1)
func (s *Stack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.l.Front().Value
}

// Push pushes an element into the stack
// O(1)
func (s *Stack) Push(v interface{}) {
	s.Lock()
	defer s.Unlock()

	s.l.Prepend(v)
	s.size++
}

// Pop returns last pushed element from the stack
// O(1)
func (s *Stack) Pop() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()

	if s.size > 0 {
		n := s.l.Front()
		ok := s.l.Remove(n)
		s.size--
		return n.Value, ok
	}

	// stack is empty
	return 0, false
}

// ToArray returns an array representation of the stack
// O(N)
func (s *Stack) ToArray() []interface{} {
	s.Lock()
	defer s.Unlock()

	return s.l.ToArray()
}
