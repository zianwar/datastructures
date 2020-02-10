package single

// ListNode is an element of a linked list.
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// List represents a singly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	head *ListNode
	tail *ListNode
	size int
}

// NewList returns an initialized singly linked list.
func NewList() *List {
	return &List{}
}

// Size returns the number of elements of list l.
// O(1)
func (l *List) Size() int {
	return l.size
}

// Append adds node to the end of the list.
// returns the appended node.
// O(1)
func (l *List) Append(v interface{}) *ListNode {
	newNode := &ListNode{Value: v}

	if l.tail != nil {
		l.tail.Next = newNode
		l.tail = newNode
	} else {
		l.tail = newNode
		l.head = newNode
	}

	l.size++
	return newNode
}

// Prepend adds node to the head of the list.
// returns the prepended node.
// O(1)
func (l *List) Prepend(v interface{}) *ListNode {
	newNode := &ListNode{Value: v}

	if l.head != nil {
		newNode.Next = l.head
		l.head = newNode
	} else {
		l.head = newNode
		l.tail = newNode
	}

	l.size++
	return newNode
}

// Remove removes all nodes that match the given value
// returns true if any node was removed, false if nothing was removed.
// O(1)
func (l *List) Remove(n *ListNode) bool {
	current := l.head

	// empty list
	if current == nil {
		return false
	}

	// remove head if necessary
	if current == n {
		current = current.Next
		l.head = current
		l.size--
		return true
	}

	// remove any non-head nodes
	removed := false
	for current != nil && current.Next != nil {
		if current.Next == n {
			current.Next = current.Next.Next
			removed = true
			l.size--
		}
		current = current.Next
	}

	return removed
}

// InsertAfter inserts a new node with value v immediately after node
// returns true if operation is successful, false otherwise.
// If node is not an element of l, the list is not modified.
// O(N)
func (l *List) InsertAfter(v interface{}, node *ListNode) bool {
	current := l.head
	newNode := &ListNode{Value: v}

	for current != nil {
		if current == node {
			newNode.Next = current.Next
			current.Next = newNode

			// if inserting after tail,
			// update tail to point to the currently inserted node
			if node == l.tail {
				l.tail = newNode
			}

			l.size++
			return true
		}
	}

	return false
}

// InsertBefore inserts a new node with value v immediately before node
// returns true if operation is successful, false otherwise.
// If node is not an element of l, the list is not modified.
// O(N)
func (l *List) InsertBefore(v interface{}, node *ListNode) bool {
	newNode := &ListNode{Value: v}

	// insert before head
	if l.head == node {
		newNode.Next = l.head
		l.head = newNode

		l.size++
		return true
	}

	// insert before non-head node
	current := l.head.Next
	prev := l.head
	for current != nil {
		if current == node {
			newNode.Next = current
			prev.Next = newNode

			l.size++
			return true
		}
		prev = current
		current = current.Next
	}

	return false
}

// Find finds an element with given value v and returns it
// O(N)
func (l *List) Find(v interface{}) *ListNode {
	current := l.head

	for current != nil {
		if current.Value == v {
			return current
		}
		current = current.Next
	}

	return nil
}

// MoveFront moves the given node to the front of the list
// O(N)
func (l *List) MoveFront(n *ListNode) bool {
	if l.Remove(n) {
		l.Prepend(n.Value)
		return true
	}
	return false
}

// MoveBack moves the given node to the back of the list
// O(N)
func (l *List) MoveBack(n *ListNode) bool {
	if l.Remove(n) {
		l.Append(n.Value)
		return true
	}

	return false
}

// ToArray returns an array representation of the values of the list l
// O(N)
func (l *List) ToArray() []interface{} {
	a := make([]interface{}, 0)
	current := l.head

	for current != nil {
		a = append(a, current.Value)
		current = current.Next
	}

	return a
}
