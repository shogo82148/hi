//go:build goexperiment.rangefunc

package itertools

import (
	"iter"
)

type Countable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Count makes an iterator that returns evenly spaced values starting with number start.
func Count[T Countable](start, step T) func(func(T) bool) {
	return func(yield func(T) bool) {
		for i := start; ; i += step {
			if !yield(i) {
				break
			}
		}
	}
}

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
