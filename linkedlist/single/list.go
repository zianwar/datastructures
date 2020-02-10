package single

import "sync"

// ListNode is an element of a linked list.
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// List represents a singly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	sync.Mutex
	front *ListNode
	tail  *ListNode
	size  int
}

// NewList returns an initialized singly linked list.
func NewList() *List {
	return &List{}
}

// Size returns the number of elements of list l.
// O(1)
func (l *List) Size() int {
	l.Lock()
	defer l.Unlock()

	return l.size
}

// Tail returns the tail of list l.
// O(1)
func (l *List) Tail() *ListNode {
	l.Lock()
	defer l.Unlock()

	return l.tail
}

// Front returns the front of list l.
// O(1)
func (l *List) Front() *ListNode {
	l.Lock()
	defer l.Unlock()

	return l.front
}

// Append adds node to the end of the list.
// returns the appended node.
// O(1)
func (l *List) Append(v interface{}) *ListNode {
	l.Lock()
	defer l.Unlock()

	newNode := &ListNode{Value: v}

	if l.tail != nil {
		l.tail.Next = newNode
		l.tail = newNode
	} else {
		l.tail = newNode
		l.front = newNode
	}

	l.size++
	return newNode
}

// Prepend adds node to the front of the list.
// returns the prepended node.
// O(1)
func (l *List) Prepend(v interface{}) *ListNode {
	l.Lock()
	defer l.Unlock()

	newNode := &ListNode{Value: v}

	if l.front != nil {
		newNode.Next = l.front
		l.front = newNode
	} else {
		l.front = newNode
		l.tail = newNode
	}

	l.size++
	return newNode
}

// Remove removes all nodes that match the given value
// returns true if any node was removed, false if nothing was removed.
// O(1) for removing front
// O(N) for removing non-front nodes
func (l *List) Remove(n *ListNode) bool {
	l.Lock()
	defer l.Unlock()

	// empty list
	if l.front == nil {
		return false
	}

	// remove front if necessary
	if l.front == n {
		l.front = l.front.Next
		l.size--
		return true
	}

	// remove any non-front nodes
	current := l.front
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
	l.Lock()
	defer l.Unlock()

	current := l.front
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
	l.Lock()
	defer l.Unlock()

	newNode := &ListNode{Value: v}

	// insert before front
	if l.front == node {
		newNode.Next = l.front
		l.front = newNode

		l.size++
		return true
	}

	// insert before non-front node
	current := l.front.Next
	prev := l.front
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
	l.Lock()
	defer l.Unlock()

	current := l.front

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
	l.Lock()
	defer l.Unlock()

	a := make([]interface{}, 0)
	current := l.front

	for current != nil {
		a = append(a, current.Value)
		current = current.Next
	}

	return a
}
