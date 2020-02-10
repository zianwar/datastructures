package queue

import (
	"sync"

	list "github.com/zianwar/datastructures/linkedlist/single"
)

// Queue represents a thread-safe queue data structure
type Queue struct {
	sync.Mutex
	l    *list.List
	size int
}

// NewQueue returns a new queue
func NewQueue() *Queue {
	return &Queue{
		l: list.NewList(),
	}
}

// Size returns the size of the queue
// O(1)
func (q *Queue) Size() int {
	q.Lock()
	defer q.Unlock()

	return q.size
}

// Peek returns the item at the front of the queue
// O(1)
func (q *Queue) Peek() interface{} {
	q.Lock()
	defer q.Unlock()

	if q.size == 0 {
		return nil
	}
	return q.l.Tail().Value
}

// Enqueue adds an item to the back of the queue
// O(1)
func (q *Queue) Enqueue(v interface{}) {
	q.Lock()
	defer q.Unlock()

	q.l.Append(v)
	q.size++
}

// Dequeue deletes and returns the item at front of the queue
// O(1)
func (q *Queue) Dequeue() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()

	if q.size > 0 {
		n := q.l.Front()
		v := n.Value
		if ok := q.l.Remove(n); !ok {
			return 0, false
		}
		q.size--
		return v, true
	}

	// Queue is empty
	return 0, false
}

// ToArray returns an array representation of the stack
// O(N)
func (q *Queue) ToArray() []interface{} {
	q.Lock()
	defer q.Unlock()

	return q.l.ToArray()
}
