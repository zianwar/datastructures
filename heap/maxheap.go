package heap

type MaxHeap struct {
	heap
}

func lessMaxHeap(items []interface{}, i, j int) bool {
	return items[i].(int) > items[j].(int)
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: *NewHeap(lessMaxHeap),
	}
}
