package stack

import "testing"

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push(2)
	s.Push(1)

	expected := []interface{}{1, 2}
	got := s.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	v, ok := s.Pop()

	if !ok {
		t.Error("operation failed")
	}

	expected := 1
	got := v
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestPeek(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	v := s.Peek()

	expected := 2
	got := v
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}

	expected2 := []interface{}{2, 1}
	got2 := s.ToArray()
	if !Equal(got2, expected2) {
		t.Errorf("expected %v got %v", expected2, got2)
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
