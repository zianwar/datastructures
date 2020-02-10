package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)

	expected := []interface{}{1, 2}
	got := q.ToArray()
	if !Equal(got, expected) {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Dequeue()
	v := q.Peek()

	expected := 2
	got := v
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	v, ok := q.Dequeue()

	if !ok {
		t.Error("operation failed")
	}

	expected := 1
	got := v
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
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
