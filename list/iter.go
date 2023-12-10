//go:build goexperiment.rangefunc

package list

import (
	"iter"
)

// Forward returns an iterator of l.
func (l *List[T]) Forward() iter.Seq2[*Element[T], T] {
	return func(yield func(*Element[T], T) bool) {
		for e := l.Front(); e != nil; e = e.Next() {
			if !yield(e, e.Value) {
				break
			}
		}
	}
}

// Reverse returns an iterator of l.
func (l *List[T]) Backward() iter.Seq2[*Element[T], T] {
	return func(yield func(*Element[T], T) bool) {
		for e := l.Back(); e != nil; e = e.Prev() {
			if !yield(e, e.Value) {
				break
			}
		}
	}
}
