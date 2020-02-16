package heap

type MinHeap struct {
	heap
}

func lessMinHeap(items []int, i, j int) bool {
	return items[i] < items[j]
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		heap: *NewHeap(lessMinHeap),
	}
}
