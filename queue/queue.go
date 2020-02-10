package queue

import (
	"sync"
)

type Queue struct {
	sync.Mutex
	items []int
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]int, 0),
	}
}

func (q *Queue) Enqueue(v int) {
	q.Lock()
	defer q.Unlock()

	q.items = append(q.items, v)
}

func (q *Queue) Dequeue() (int, bool) {
	q.Lock()
	defer q.Unlock()

	if len(q.items) > 0 {
		v := q.items[0]
		q.items = q.items[0:]
		return v, true
	}

	// Queue is empty
	return 0, false
}
