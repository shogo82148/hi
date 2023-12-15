//go:build goexperiment.rangefunc

package it

import (
	"iter"
)

// SliceIter returns an iterator for the slice.
func SliceIter[S ~[]E, E any](x S) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i, v := range x {
			if !yield(i, v) {
				break
			}
		}
	}
}

// SliceValues returns an iterator for the slice.
func SliceValues[S ~[]E, E any](x S) func(func(E) bool) {
	return func(yield func(E) bool) {
		for _, v := range x {
			if !yield(v) {
				break
			}
		}
	}
}

// MapIter returns an iterator for the map.
func MapIter[M ~map[K]V, K comparable, V any](x M) func(func(K, V) bool) {
	return func(yield func(K, V) bool) {
		for k, v := range x {
			if !yield(k, v) {
				break
			}
		}
	}
}

// MapKeys returns an iterator for the map.
func MapKeys[M ~map[K]V, K comparable, V any](x M) func(func(K) bool) {
	return func(yield func(K) bool) {
		for k := range x {
			if !yield(k) {
				break
			}
		}
	}
}

// MapValues returns an iterator for the map.
func MapValues[M ~map[K]V, K comparable, V any](x M) func(func(V) bool) {
	return func(yield func(V) bool) {
		for _, v := range x {
			if !yield(v) {
				break
			}
		}
	}
}

// ChanValues returns an iterator for the channel.
func ChanValues[T any](x <-chan T) func(func(T) bool) {
	return func(yield func(T) bool) {
		for v := range x {
			if !yield(v) {
				break
			}
		}
	}
}

func Range(n int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			if !yield(i) {
				break
			}
		}
	}
}

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

// Filter returns an iterator that filters elements from data returning only those
// that have a corresponding element in selectors that evaluates to True.
func Filter[T any](seq iter.Seq[T], filter func(T) bool) func(func(T) bool) {
	return func(yield func(T) bool) {
		for v := range seq {
			if filter(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// Filter2 returns an iterator that filters elements from data returning only those
// that have a corresponding element in selectors that evaluates to True.
func Filter2[K, V any](seq iter.Seq2[K, V], filter func(K, V) bool) func(func(K, V) bool) {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if filter(k, v) {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}
