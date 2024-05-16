package linkedlist

import "github.com/TannerKvarfordt/ubiquity/err"

type node[T any] struct {
	data T
	next *node[T]
}

// SinglyLinkedList dynamically stores data in non-sequential memory.
type SinglyLinkedList[T any] struct {
	head *node[T]
}

// NewSinglyLinkedList creates a new SinglyLinkedList containing
// any provided data.
func NewSinglyLinkedList[T any](data ...T) *SinglyLinkedList[T] {
	l := SinglyLinkedList[T]{}
	for _, d := range data {
		l.Append(d)
	}
	return &l
}

// Len returns the number of elements in the list.
func (l *SinglyLinkedList[_]) Len() int {
	len := 0
	for curr := l.head; curr != nil; curr = curr.next {
		len++
	}
	return len
}

// ForEach iterates over all items in the list, calling
// fn on each.
func (l *SinglyLinkedList[T]) ForEach(fn func(data T)) {
	if fn == nil {
		return
	}
	for curr := l.head; curr != nil; curr = curr.next {
		fn(curr.data)
	}
}

// Append appends a node to the list containing the
// given data.
// TODO: accept variadic args
func (l *SinglyLinkedList[T]) Append(data T) {
	if l.head == nil {
		l.head = &node[T]{data: data}
		return
	}

	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = &node[T]{data: data, next: nil}
}

// Prepend prepends a node to the list containing the
// given data.
// TODO: accept variadic args
func (l *SinglyLinkedList[T]) Prepend(data T) {
	next := l.head
	l.head = &node[T]{data: data, next: next}
}

// Insert inserts the data at the given index.
// Indices for any data beyond the insertion point
// are increased by 1.
//
// Valid indices for this function are in the range [0, l.Len()].
//
// Returns an error if the index is out of bounds.
func (l *SinglyLinkedList[T]) Insert(index int, data T) error {
	if index < 0 || index > l.Len() {
		return &err.IndexOutOfRange{Idx: index}
	}

	if index == l.Len() {
		l.Append(data)
		return nil
	}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	tmp := &node[T]{data: curr.data, next: curr.next}
	curr.data = data
	curr.next = tmp
	return nil
}

// SetData sets the data value at the provided index.
// If the index is out of range, an error is returned.
func (l *SinglyLinkedList[T]) SetData(index int, data T) error {
	if index < 0 || index >= l.Len() {
		return &err.IndexOutOfRange{Idx: index}
	}

	curr := l.head
	for i := 0; i != index; i++ {
		curr = curr.next
	}
	curr.data = data
	return nil
}

// GetData returns a copy of the data at the given index in
// the list. If the index is out of range, a zero-value and
// an error will be returned.
func (l *SinglyLinkedList[T]) GetData(index int) (T, error) {
	if index < 0 || index >= l.Len() {
		return *new(T), &err.IndexOutOfRange{Idx: index}
	}

	curr := l.head
	for i := 0; i != index; i++ {
		curr = curr.next
	}
	return curr.data, nil
}

// Delete removes from the list the item at the given index.
// The value at the given index is returned. If the index
// is out of range, a zero-value and an error will be returned.
func (l *SinglyLinkedList[T]) Delete(index int) (T, error) {
	if index < 0 || index >= l.Len() {
		return *new(T), &err.IndexOutOfRange{Idx: index}
	}

	if index == 0 {
		rv := l.head.data
		l.head = l.head.next
		return rv, nil
	}

	curr := l.head
	for i := 0; i != index-1; i++ {
		curr = curr.next
	}
	rv := curr.next.data
	curr.next = curr.next.next
	return rv, nil
}

// TODO: func (l *SinglyLinkedList[T]) Find(data T) (indices []int, err error)
// TODO: func (l *SinglyLinkedList[T]) Sort()
// TODO: func Merge[T any](a, b SinglyLinkedList[T]) *SinglyLinkedList[T]
// TODO: func Equal[T any](a, b SinglyLinkedList[T]) bool
