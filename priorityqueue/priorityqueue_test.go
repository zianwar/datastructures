package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	pq.Enqueue(1)
	pq.Enqueue(10)
	pq.Enqueue(5)
	pq.Enqueue(7)

	assert.EqualValues(t, []interface{}{10, 7, 1, 5}, pq.ToArray())
	x, _ := pq.Dequeue()
	assert.Equal(t, 10, x.(int))
	assert.EqualValues(t, []interface{}{7, 5, 1}, pq.ToArray())
}
