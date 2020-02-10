package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)

	got, _ := q.Dequeue()
	expected := 1

	if got != expected {
		t.Errorf("got %v expected %v", got, expected)
	}
}
