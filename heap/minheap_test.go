package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()
	h.Push(5)
	h.Push(6)
	h.Push(7)
	h.Push(10)
	h.Push(1)

	assert.Equal(t, 5, len(h.items), "size")
	assert.Equal(t, 1, h.Pop().(int), "pop")
	assert.Equal(t, 4, len(h.items), "size decreased after Pop")
	assert.Equal(t, 5, h.Pop().(int), "pop")
	assert.Equal(t, 6, h.Peek().(int), "return min")
}
