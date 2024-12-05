// itertools contains functions for creating iterators.
// it is inspired by Python's itertools.

package it

import (
	"iter"

	"github.com/shogo82148/hi/tuple"
)

// Cycle makes an iterator returning elements from the iterable and saving a copy of each.
func Cycle[T any](seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		var saved []T
		for v := range seq {
			if !yield(v) {
				return
			}
			saved = append(saved, v)
		}
		for len(saved) > 0 {
			for _, v := range saved {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Cycle2 makes an iterator returning elements from the iterable and saving a copy of each.
func Cycle2[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var saved []tuple.Tuple2[K, V]
		for k, v := range seq {
			if !yield(k, v) {
				return
			}
			saved = append(saved, tuple.New2(k, v))
		}
		for len(saved) > 0 {
			for _, t := range saved {
				if !yield(t.V1, t.V2) {
					return
				}
			}
		}
	}
}

// Repeat makes an iterator that returns object over and over again.
func Repeat[T any](v T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if !yield(v) {
				break
			}
		}
	}
}

// RepeatN makes an iterator that returns object over and over again.
func RepeatN[T any](v T, n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for range n {
			if !yield(v) {
				break
			}
		}
	}
}

// Accumulate makes an iterator that returns accumulated sums,
// or accumulated results of other binary functions (specified via the optional func argument).
func Accumulate[T any](seq iter.Seq[T], f func(T, T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		var total T
		initialized := false
		for v := range seq {
			if !initialized {
				total = v
				initialized = true
			} else {
				total = f(total, v)
			}
			if !yield(total) {
				break
			}
		}
	}
}

// Chain makes an iterator that returns elements from the first iterable until it is exhausted,
// then proceeds to the next iterable, until all of the iterables are exhausted.
func Chain[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Chain2 makes an iterator that returns elements from the first iterable until it is exhausted,
// then proceeds to the next iterable, until all of the iterables are exhausted.
func Chain2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// ChainFromIterables makes an iterator that returns elements from the first iterable until it is exhausted,
// then proceeds to the next iterable, until all of the iterables are exhausted.
func ChainFromIterables[T any](seqs iter.Seq[iter.Seq[T]]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for seq := range seqs {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// ChainFromIterables2 makes an iterator that returns elements from the first iterable until it is exhausted,
// then proceeds to the next iterable, until all of the iterables are exhausted.
func ChainFromIterables2[K, V any](seqs iter.Seq[iter.Seq2[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for seq := range seqs {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Compress makes an iterator that filters elements from data returning only those
// that have a corresponding element in selectors that evaluates to true.
func Compress[T any](data iter.Seq[T], selectors iter.Seq[bool]) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(selectors)
		defer stop()

		for v := range data {
			s, ok := next()
			if !ok {
				break
			}
			if s {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// Compress2 makes an iterator that filters elements from data returning only those
// that have a corresponding element in selectors that evaluates to true.
func Compress2[K, V any](data iter.Seq2[K, V], selectors iter.Seq[bool]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull(selectors)
		defer stop()

		for k, v := range data {
			s, ok := next()
			if !ok {
				break
			}
			if s {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// DropWhile makes an iterator that drops elements from the iterable as long as the predicate is true;
// afterwards, returns every element.
func DropWhile[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		dropping := true
		for v := range seq {
			if dropping {
				if f(v) {
					continue
				}
				dropping = false
			}
			if !yield(v) {
				break
			}
		}
	}
}

// DropWhile2 makes an iterator that drops elements from the iterable as long as the predicate is true;
// afterwards, returns every element.
func DropWhile2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		dropping := true
		for k, v := range seq {
			if dropping {
				if f(k, v) {
					continue
				}
				dropping = false
			}
			if !yield(k, v) {
				break
			}
		}
	}
}

// FilterFalse makes an iterator that filters elements from data returning only those
// that have a corresponding element in selectors that evaluates to False.
func FilterFalse[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !f(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// FilterFalse2 makes an iterator that filters elements from data returning only those
// that have a corresponding element in selectors that evaluates to False.
func FilterFalse2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !f(k, v) {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// Slice makes an iterator that returns selected elements from the iterable.
// If start is non-zero, then elements from the iterable are skipped until start is reached.
// Afterward, elements are returned consecutively unless step is set higher than one which results in items being skipped.
// If stop is negative, then iteration continues until the iterator is exhausted, if at all; otherwise, it stops at the specified position.
//
// start must be non-negative, and step must be positive.
func Slice[T any](seq iter.Seq[T], start, stop, step int) iter.Seq[T] {
	if start < 0 {
		panic("it: start must be non-negative")
	}
	if step <= 0 {
		panic("it: step must be positive")
	}

	// special cases
	if stop < 0 && step == 1 {
		return func(yield func(T) bool) {
			var i int
			for v := range seq {
				if i >= start {
					if !yield(v) {
						break
					}
				}
				i++
			}
		}
	}
	if stop < 0 {
		return func(yield func(T) bool) {
			var i int
			for v := range seq {
				if i >= start && (i-start)%step == 0 {
					if !yield(v) {
						break
					}
				}
				i++
			}
		}
	}
	if step == 1 {
		return func(yield func(T) bool) {
			var i int
			for v := range seq {
				if i >= stop {
					break
				}
				if i >= start {
					if !yield(v) {
						break
					}
				}
				i++
			}
		}
	}

	// general case
	return func(yield func(T) bool) {
		var i int
		for v := range seq {
			if i >= stop {
				break
			}
			if i >= start && (i-start)%step == 0 {
				if !yield(v) {
					break
				}
			}
			i++
		}
	}
}

func Slice2[K, V any](seq iter.Seq2[K, V], start, stop, step int) iter.Seq2[K, V] {
	if start < 0 {
		panic("it: start must be non-negative")
	}
	if step <= 0 {
		panic("it: step must be positive")
	}

	// special cases
	if stop < 0 && step == 1 {
		return func(yield func(K, V) bool) {
			var i int
			for k, v := range seq {
				if i >= start {
					if !yield(k, v) {
						break
					}
				}
				i++
			}
		}
	}
	if stop < 0 {
		return func(yield func(K, V) bool) {
			var i int
			for k, v := range seq {
				if i >= start && (i-start)%step == 0 {
					if !yield(k, v) {
						break
					}
				}
				i++
			}
		}
	}
	if step == 1 {
		return func(yield func(K, V) bool) {
			var i int
			for k, v := range seq {
				if i >= stop {
					break
				}
				if i >= start {
					if !yield(k, v) {
						break
					}
				}
				i++
			}
		}
	}

	// general case
	return func(yield func(K, V) bool) {
		var i int
		for k, v := range seq {
			if i >= stop {
				break
			}
			if i >= start && (i-start)%step == 0 {
				if !yield(k, v) {
					break
				}
			}
			i++
		}
	}
}

// Pairwise makes an iterator that returns tuples of the previous and current items of seq.
func Pairwise[T any](seq iter.Seq[T]) iter.Seq[tuple.Tuple2[T, T]] {
	return func(yield func(tuple.Tuple2[T, T]) bool) {
		var prev T
		var initialized bool
		for v := range seq {
			if initialized {
				if !yield(tuple.New2(prev, v)) {
					break
				}
			}
			prev = v
			initialized = true
		}
	}
}

// Pairwise2 makes an iterator that returns tuples of the previous and current items of seq.
func Pairwise2[K, V any](seq iter.Seq2[K, V]) iter.Seq2[tuple.Tuple2[K, K], tuple.Tuple2[V, V]] {
	return func(yield func(tuple.Tuple2[K, K], tuple.Tuple2[V, V]) bool) {
		var prevK K
		var prevV V
		var initialized bool
		for k, v := range seq {
			if initialized {
				if !yield(tuple.New2(prevK, k), tuple.New2(prevV, v)) {
					break
				}
			}
			prevK, prevV = k, v
			initialized = true
		}
	}
}

// TakeWhile makes an iterator that returns elements from the iterable as long as f returns true.
func TakeWhile[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !f(v) {
				break
			}
			if !yield(v) {
				break
			}
		}
	}
}

// TakeWhile2 makes an iterator that returns elements from the iterable as long as f returns true.
func TakeWhile2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !f(k, v) {
				break
			}
			if !yield(k, v) {
				break
			}
		}
	}
}

// Tee makes n iterators from seq.
func Tee[T any](seq iter.Seq[T], n int) []iter.Seq[T] {
	if n < 0 {
		panic("it: n must be non-negative")
	}
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []iter.Seq[T]{seq}
	}

	que := newList[T]()
	elements := make([]*element[T], n)
	lifeFirst := n
	first := que.last

	life := n
	next, stop := iter.Pull(seq)

	ret := make([]iter.Seq[T], n)
	for i := range n {
		ret[i] = func(yield func(T) bool) {
			defer func() {
				life--
				if life == 0 {
					// all iterators are stopped.
					stop()
				}
			}()

			for {
				e := elements[i]
				if e == nil {
					e = first

					// garbage collector now can collect the first.
					lifeFirst--
					if lifeFirst == 0 {
						first = nil
					}
				}
				if e == que.last {
					v, ok := next()
					if !ok {
						break
					}
					que.pushBack(v)
					e = que.tail
					if first == que.last {
						first = e
					}
				}

				if !yield(e.value) {
					break
				}
				elements[i] = e.next
			}
		}
	}
	return ret
}

// Tee2 makes n iterators from seq.
func Tee2[K, V any](seq iter.Seq2[K, V], n int) []iter.Seq2[K, V] {
	if n < 0 {
		panic("it: n must be non-negative")
	}
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []iter.Seq2[K, V]{seq}
	}

	que := newList[tuple.Tuple2[K, V]]()
	elements := make([]*element[tuple.Tuple2[K, V]], n)
	lifeFirst := n
	first := que.last

	life := n
	next, stop := iter.Pull2(seq)

	ret := make([]iter.Seq2[K, V], n)
	for i := range n {
		ret[i] = func(yield func(K, V) bool) {
			defer func() {
				life--
				if life == 0 {
					// all iterators are stopped.
					stop()
				}
			}()

			for {
				e := elements[i]
				if e == nil {
					e = first

					// garbage collector now can collect the first.
					lifeFirst--
					if lifeFirst == 0 {
						first = nil
					}
				}
				if e == que.last {
					k, v, ok := next()
					if !ok {
						break
					}
					que.pushBack(tuple.New2(k, v))
					e = que.tail
					if first == que.last {
						first = e
					}
				}

				if !yield(e.value.V1, e.value.V2) {
					break
				}
				elements[i] = e.next
			}
		}
	}
	return ret
}
