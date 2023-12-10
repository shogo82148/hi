//go:build goexperiment.rangefunc

package itertools

import (
	"iter"
)

// Append appends all elements of seq to s and returns the resulting slice.
func Append[S ~[]E, E any](s S, seq iter.Seq[E]) S {
	for v := range seq {
		s = append(s, v)
	}
	return s
}

// KeyValues makes an iterator that returns keys and values from two iterators.
func KeyValues[K, V any](keys iter.Seq[K], values iter.Seq[V]) func(func(K, V) bool) {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull(values)
		defer stop()

		for k := range keys {
			v, ok := next()
			if !ok {
				break
			}
			if !yield(k, v) {
				break
			}
		}
	}
}

// Keys makes an iterator that returns keys from an iterator.
func Keys[K, V any](seq iter.Seq2[K, V]) func(func(K) bool) {
	return func(yield func(K) bool) {
		for k, _ := range seq {
			if !yield(k) {
				break
			}
		}
	}
}

// Values makes an iterator that returns values from an iterator.
func Values[K, V any](seq iter.Seq2[K, V]) func(func(V) bool) {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				break
			}
		}
	}
}
