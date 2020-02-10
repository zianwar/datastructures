package single

import (
	"testing"
)

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

func TestAppend(t *testing.T) {
	l := NewList()
	l.Append(1)
	l.Append(2)

	expected := []interface{}{1, 2}
	got := l.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestPrepend(t *testing.T) {
	l := NewList()
	l.Prepend(1)
	l.Prepend(2)

	expected := []interface{}{2, 1}
	got := l.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestRemove(t *testing.T) {
	l := NewList()
	n := l.Append(1)
	l.Append(2)
	l.Remove(n)

	expected := []interface{}{2}
	got := l.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestFind(t *testing.T) {
	l := NewList()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	n := l.Find(2)

	expected := 2
	got := n.Value
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestInsertAfter(t *testing.T) {
	l := NewList()
	n := l.Append(1)
	l.Append(3)

	ok := l.InsertAfter(2, n)

	expected := []interface{}{1, 2, 3}
	got := l.ToArray()
	if !ok || !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestInsertBefore(t *testing.T) {
	l := NewList()
	l.Append(1)
	n := l.Append(3)

	if ok := l.InsertBefore(2, n); !ok {
		t.Error("InsertBefore failed")
	}

	expected := []interface{}{1, 2, 3}
	got := l.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestMoveFront(t *testing.T) {
	l := NewList()
	l.Append(2)
	l.Append(3)
	l.Append(4)
	n := l.Append(1)

	if ok := l.MoveFront(n); !ok {
		t.Error("operation failed")
	}

	expected := []interface{}{1, 2, 3, 4}
	got := l.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestMoveBack(t *testing.T) {
	l := NewList()
	n := l.Append(4)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if ok := l.MoveBack(n); !ok {
		t.Error("operation failed")
	}

	expected := []interface{}{1, 2, 3, 4}
	got := l.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}
