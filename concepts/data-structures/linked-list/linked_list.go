package linkedlist

// Node stores one value and a link to the next node.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// FromSlice returns a list containing values in their slice order.
func FromSlice[T any](values []T) *Node[T] {
	dummy := &Node[T]{}
	tail := dummy
	for _, value := range values {
		tail.Next = &Node[T]{Value: value}
		tail = tail.Next
	}
	return dummy.Next
}

// ToSlice returns the values encountered from head to nil.
func ToSlice[T any](head *Node[T]) []T {
	values := make([]T, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Value)
	}
	return values
}

// Reverse reverses a list in place and returns its new head.
func Reverse[T any](head *Node[T]) *Node[T] {
	var previous *Node[T]
	for current := head; current != nil; {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
