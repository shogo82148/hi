//go:build goexperiment.rangefunc

package it

// minimum implementation of single-linked list.

type element[T any] struct {
	value T
	next  *element[T]
}

type list[T any] struct {
	// last is the sentinel element of the end of the list.
	last *element[T]

	// tail is the last element of the list.
	// if the list is empty, tail == last.
	tail *element[T]
}

func newList[T any]() *list[T] {
	last := new(element[T])
	last.next = last

	return &list[T]{
		last: last,
		tail: last,
	}
}

func (l *list[T]) pushBack(v T) {
	e := &element[T]{
		value: v,
		next:  l.last,
	}
	if l.tail != l.last {
		l.tail.next = e
	}
	l.tail = e
}
