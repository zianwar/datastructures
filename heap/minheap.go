package heap

type MinHeap struct {
	heap
}

func lessMinHeap(items []interface{}, i, j int) bool {
	return items[i].(int) < items[j].(int)
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		heap: *NewHeap(lessMinHeap),
	}
}
