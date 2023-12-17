//go:build goexperiment.rangefunc

// itertools contains functions for creating iterators.
// it is inspired by Python's itertools.

package it

import (
	"iter"
)

// Cycle makes an iterator returning elements from the iterable and saving a copy of each.
func Cycle[T any](seq iter.Seq[T]) func(func(T) bool) {
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

// Repeat makes an iterator that returns object over and over again.
func Repeat[T any](v T) func(func(T) bool) {
	return func(yield func(T) bool) {
		for {
			if !yield(v) {
				break
			}
		}
	}
}

// RepeatN makes an iterator that returns object over and over again.
func RepeatN[T any](v T, n int) func(func(T) bool) {
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
func Accumulate[T any](seq iter.Seq[T], f func(T, T) T) func(func(T) bool) {
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
func Chain[T any](seqs ...iter.Seq[T]) func(func(T) bool) {
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

// ChainFromIterables makes an iterator that returns elements from the first iterable until it is exhausted,
// then proceeds to the next iterable, until all of the iterables are exhausted.
func ChainFromIterables[T any](seqs iter.Seq[iter.Seq[T]]) func(func(T) bool) {
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

// Compress makes an iterator that filters elements from data returning only those
// that have a corresponding element in selectors that evaluates to True.
func Compress[T any](data iter.Seq[T], selectors iter.Seq[bool]) func(func(T) bool) {
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

// DropWhile makes an iterator that drops elements from the iterable as long as the predicate is true;
// afterwards, returns every element.
func DropWhile[T any](f func(T) bool, seq iter.Seq[T]) func(func(T) bool) {
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

// FilterFalse makes an iterator that filters elements from data returning only those
// that have a corresponding element in selectors that evaluates to False.
func FilterFalse[T any](seq iter.Seq[T], f func(T) bool) func(func(T) bool) {
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

// Slice makes an iterator that returns selected elements from the iterable.
// If start is non-zero, then elements from the iterable are skipped until start is reached.
// Afterward, elements are returned consecutively unless step is set higher than one which results in items being skipped.
// If stop is negative, then iteration continues until the iterator is exhausted, if at all; otherwise, it stops at the specified position.
//
// start must be non-negative, and step must be positive.
func Slice[T any](seq iter.Seq[T], start, stop, step int) func(func(T) bool) {
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

func Slice2[K, V any](seq iter.Seq2[K, V], start, stop, step int) func(func(K, V) bool) {
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
