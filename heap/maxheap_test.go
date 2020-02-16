package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()
	h.Push(5)
	h.Push(6)
	h.Push(7)
	h.Push(10)
	h.Push(1)

	assert.Equal(t, 5, h.Len(), "size")
	assert.Equal(t, 10, *h.Pop(), "pop")
	assert.Equal(t, 4, h.Len(), "size decreased after Pop")
	assert.Equal(t, 7, *h.Pop(), "pop")
	assert.Equal(t, 6, *h.Peek(), "return min")
}
