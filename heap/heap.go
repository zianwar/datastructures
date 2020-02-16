package heap

type comparator func(items []interface{}, i, j int) bool

type heap struct {
	items []interface{}
	less  comparator
}

func NewHeap(less comparator) *heap {
	return &heap{
		items: make([]interface{}, 0),
		less:  less,
	}
}

func (h *heap) Len() int {
	return len(h.items)
}

func (h *heap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// Push inserts an item into the min-heap
//
// We always start by inserting the element at the bottom.
// We insert at the rightmost spot so as to maintain the complete tree property.
//
// Then, we "fix" the tree by swapping the new element with its parent,
// until we find an appropriate spot for the element.
// We essentially bubble up the minimum element.
// O(log(n))
func (h *heap) Push(v interface{}) {
	h.items = append(h.items, v)
	h.heapifyUp()
}

// Peek returns the min without removing it
// O(1)
func (h *heap) Peek() interface{} {
	if h.Len() > 0 {
		return h.items[0]
	}
	return nil
}

// Pop return the min item of the heap which is always at the top.
//
// The trickier part is how to remove it:
// 1. First, we remove the minimum element and swap it with the last element
//    in the heap (the bottommost, rightmost element).
// 2. Then, we bubble down this element, swapping it with one of its children
//    until the min-heap property is restored.
// 3. Do we swap it with the left child or the right child?
//    That depends on their values. There's no inherent ordering between
//    the left and right.
// O(log(n))
func (h *heap) Pop() interface{} {
	if h.Len() > 0 {
		// save the min at the front and move the last item to the front.
		min := h.items[0]
		h.items[0] = h.items[h.Len()-1]

		// delete the last item
		h.items = h.items[:h.Len()-1]

		// heapify to correctly place the item at the front
		h.heapifyDown()
		return min
	}
	return nil
}

func (h *heap) heapifyUp() {
	if h.Len() <= 1 {
		return
	}

	// start from the end of the array going backwards
	// swap each item if its value is less than its parent
	i := h.Len() - 1
	parentIndex := i / 2

	for i > 0 {
		if h.less(h.items, i, parentIndex) {
			h.Swap(i, parentIndex)
		}

		i = parentIndex
		parentIndex = i / 2
	}
}

func (h *heap) heapifyDown() {
	if h.Len() <= 1 {
		return
	}

	// start from the begining of the array, and swap with smallest of children
	// move i to the swapped children and continue
	i := 0

	for i < h.Len() {
		left := 2*i + 1
		right := 2*i + 2

		smallest := i

		if left < h.Len() && h.less(h.items, left, smallest) {
			smallest = left
		}

		if right < h.Len() && h.less(h.items, right, smallest) {
			smallest = right
		}

		if i != smallest {
			h.Swap(i, smallest)
			i = smallest
		} else if left < h.Len() {
			i = left
		} else {
			i = right
		}
	}
}

func (h *heap) ToArray() []interface{} {
	return h.items
}
