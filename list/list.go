// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package list implements a generic doubly linked list.
// It is re-implementation of [container/list] with generics.
package list

//go:generate ./generate-zip.pl
//go:generate ./generate-unzip.pl

type Element[T any] struct {
	Value T

	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element[T]

	// The list to which this element belongs.
	list *List[T]
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type List[T any] struct {
	root Element[T] // sentinel list element, only &root, root.prev, and root.next are used
	len  int        // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List[T]) Init() *List[T] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New returns an initialized list.
func New[T any]() *List[T] { return new(List[T]).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List[T]) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *List[T]) Front() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero List value.
func (l *List[T]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List[T]) insert(e, at *Element[T]) *Element[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List[T]) insertValue(v T, at *Element[T]) *Element[T] {
	return l.insert(&Element[T]{Value: v}, at)
}

// remove removes e from its list, decrements l.len, and returns e.
func (l *List[T]) remove(e *Element[T]) *Element[T] {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}

// move moves e to next to at and returns e.
func (l *List[T]) move(e, at *Element[T]) *Element[T] {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List[T]) Remove(e *Element[T]) T {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List[T]) PushFront(v T) *Element[T] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List[T]) PushBack(v T) *Element[T] {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToFront(e *Element[T]) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToBack(e *Element[T]) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[T]) PushBackList(other *List[T]) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[T]) PushFrontList(other *List[T]) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// generic functions

// Filter iterates over elements of collection, returning a list of all elements predicate returns true for.
func Filter[T any](l *List[T], f func(index int, value T) bool) *List[T] {
	var ret List[T]
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if f(i, e.Value) {
			ret.PushBack(e.Value)
		}
	}
	return &ret
}

// FilterFalse iterates over elements of collection, returning a list of all elements predicate returns false for.
func FilterFalse[T any](l *List[T], f func(index int, value T) bool) *List[T] {
	var ret List[T]
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if !f(i, e.Value) {
			ret.PushBack(e.Value)
		}
	}
	return &ret
}

// Count counts the number of elements in the collection that compare equal to value.
func Count[T comparable](l *List[T], value T) int {
	var count int
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if e.Value == value {
			count++
		}
	}
	return count
}

// CountBy counts the number of elements that counter returns true.
func CountBy[T any](l *List[T], f func(index int, value T) bool) int {
	var count int
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if f(i, e.Value) {
			count++
		}
	}
	return count
}

// Any returns whether l has value at least one.
func Any[T comparable](l *List[T], value T) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return true
		}
	}
	return false
}

// AnyBy returns whether l has an element for that f returns true.
func AnyBy[T any](l *List[T], f func(index int, value T) bool) bool {
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if f(i, e.Value) {
			return true
		}
	}
	return false
}

// All returns whether all elements of l are value.
func All[T comparable](l *List[T], value T) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value != value {
			return false
		}
	}
	return true
}

// AllBy returns whether f returns true for all elements in l.
func AllBy[T any](l *List[T], f func(int, T) bool) bool {
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if !f(i, e.Value) {
			return false
		}
	}
	return true
}

// Map
func Map[T, U any](l *List[T], f func(int, T) U) *List[U] {
	var ret List[U]
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		v := f(i, e.Value)
		ret.PushBack(v)
	}
	return &ret
}

// Reduce reduces l.
func Reduce[T, U any](l *List[T], f func(i int, acc U, item T) U, init U) U {
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		init = f(i, init, e.Value)
	}
	return init
}
