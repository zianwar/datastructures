package priorityqueue

import (
	"sync"

	"github.com/zianwar/datastructures/heap"
)

// PriorityQueue represents a thread-safe priority queue data structure
// using a max heap, high priority means high value for this implementation.
type PriorityQueue struct {
	heap heap.MaxHeap
	size int
	mu   sync.Mutex
}

// NewPriorityQueue returns a new priority queue.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		heap: *heap.NewMaxHeap(),
		mu:   sync.Mutex{},
	}
}

// Size returns the size of the priority queue.
// O(1)
func (pq *PriorityQueue) Size() int {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	return pq.size
}

// Peek returns the item at the front (highest-priority).
// O(1)
func (pq *PriorityQueue) Peek() interface{} {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if pq.size == 0 {
		return nil
	}
	return pq.heap.Peek()
}

// Enqueue adds an item to the priority queue.
// O(log(n))
func (pq *PriorityQueue) Enqueue(v interface{}) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	pq.heap.Push(v)
	pq.size++
}

// Dequeue deletes and returns the item at the front (highest priority).
// O(log(n))
func (pq *PriorityQueue) Dequeue() (interface{}, bool) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	front := pq.heap.Pop()
	if front != nil {
		return front, true
	}

	// is empty
	return 0, false
}

// ToArray returns an array representation of the priority queue.
// O(N)
func (pq *PriorityQueue) ToArray() []interface{} {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	return pq.heap.ToArray()
}
