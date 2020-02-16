package heap

type MaxHeap struct {
	heap
}

func lessMaxHeap(items []int, i, j int) bool {
	return items[i] > items[j]
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: *NewHeap(lessMaxHeap),
	}
}
