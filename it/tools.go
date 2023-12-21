//go:build goexperiment.rangefunc

package it

import (
	"cmp"
	"iter"
	"runtime"

	"github.com/shogo82148/hi/optional"
	"github.com/shogo82148/hi/tuple"
	"github.com/shogo82148/hi"
)

//go:generate ./generate-zip.pl
//go:generate ./generate-unzip.pl

// pullSeq is a pull-based iterator.
type pullSeq[T any] struct {
	next func() (T, bool)
	stop func()
}

func newPullSeq[T any](seq iter.Seq[T]) *pullSeq[T] {
	next, stop := iter.Pull(seq)
	pull := &pullSeq[T]{next, stop}

	runtime.SetFinalizer(pull, func(pull *pullSeq[T]) {
		pull.stop()
	})
	return pull
}

type pullSeq2[K, V any] struct {
	next func() (K, V, bool)
	stop func()
}

func newPullSeq2[K, V any](seq iter.Seq2[K, V]) *pullSeq2[K, V] {
	next, stop := iter.Pull2(seq)
	pull := &pullSeq2[K, V]{next, stop}

	runtime.SetFinalizer(pull, func(pull *pullSeq2[K, V]) {
		pull.stop()
	})
	return pull
}

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

// Range returns an iterator for the range.
func Range(n int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			if !yield(i) {
				break
			}
		}
	}
}

// Enumerate returns an iterator that returns (index, value) tuples from seq.
func Enumerate[T any](seq iter.Seq[T]) func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		var i int
		for v := range seq {
			if !yield(i, v) {
				break
			}
			i++
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

// KeyValues makes an iterator that returns (key, value) tuples from an iterator.
func KeyValues[K, V any](seq iter.Seq[tuple.Tuple2[K, V]]) func(func(K, V) bool) {
	return func(yield func(K, V) bool) {
		for kv := range seq {
			if !yield(kv.V1, kv.V2) {
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

// Map calls mapper for each element of seq and returns the result in a iterator.
func Map[T1, T2 any](seq iter.Seq[T1], mapper func(T1) T2) func(func(T2) bool) {
	return func(yield func(T2) bool) {
		for v := range seq {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

// Map2 calls mapper for each element of seq and returns the result in a iterator.
func Map2[K1, V1, K2, V2 any](seq iter.Seq2[K1, V1], mapper func(K1, V1) (K2, V2)) func(func(K2, V2) bool) {
	return func(yield func(K2, V2) bool) {
		for k, v := range seq {
			if !yield(mapper(k, v)) {
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

// Zip is similar to Zip2, but it returns [iter.Seq2].
func Zip[K, V any](keys iter.Seq[K], values iter.Seq[V]) func(func(K, V) bool) {
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

// ZipLongest make an iterator that aggregates elements from each of the iterables.
// If the iterables are of uneven length, missing values are filled-in with zero-values.
// Iteration continues until the longest iterable is exhausted.
func ZipLongest[K, V any](keys iter.Seq[K], values iter.Seq[V]) func(func(K, V) bool) {
	return func(yield func(K, V) bool) {
		nextK, stopK := iter.Pull(keys)
		defer stopK()
		nextV, stopV := iter.Pull(values)
		defer stopV()

		for {
			k, ok1 := nextK()
			v, ok2 := nextV()
			if !ok1 && !ok2 {
				break
			}
			if !ok1 {
				k = hi.Zero[K]()
			}
			if !ok2 {
				v = hi.Zero[V]()
			}
			if !yield(k, v) {
				break
			}
		}
	}
}

// Unzip is similar to Unzip2, but it accepts [iter.Seq2].
func Unzip[K, V any](seq iter.Seq2[K, V]) (func(func(K) bool), func(func(V) bool)) {
	s := Tee2(seq, 2)

	return func(yield func(K) bool) {
			for k, _ := range s[0] {
				if !yield(k) {
					break
				}
			}
		}, func(yield func(V) bool) {
			for _, v := range s[1] {
				if !yield(v) {
					break
				}
			}
		}
}

// GroupBy returns a map of slices, where each slice contains elements of seq grouped by the result of key.
func GroupBy[K comparable, V any](seq iter.Seq[V], key func(V) K) map[K][]V {
	m := make(map[K][]V)
	for v := range seq {
		k := key(v)
		m[k] = append(m[k], v)
	}
	return m
}

// Chunk returns an iterator that returns slices of length size from seq.
func Chunk[T any](seq iter.Seq[T], size int) func(func([]T) bool) {
	return func(yield func([]T) bool) {
		chunk := make([]T, 0, size)
		for v := range seq {
			chunk = append(chunk, v)
			if len(chunk) == size {
				if !yield(chunk) {
					break
				}
				chunk = chunk[:0]
			}
		}
		if len(chunk) > 0 {
			yield(chunk)
		}
	}
}

// Count counts the number of elements in the collection that compare equal to value.
func Count[T comparable](seq iter.Seq[T], value T) int {
	var count int
	for v := range seq {
		if v == value {
			count++
		}
	}
	return count
}

// CountBy counts the number of elements that counter returns true.
func CountBy[T any](seq iter.Seq[T], counter func(T) bool) int {
	var count int
	for v := range seq {
		if counter(v) {
			count++
		}
	}
	return count
}

// Any returns whether seq has value at least one.
func Any[T comparable](seq iter.Seq[T], value T) bool {
	for v := range seq {
		if v == value {
			return true
		}
	}
	return false
}

// Any2 returns whether seq has key-value pair at least one.
func Any2[K, V comparable](seq iter.Seq2[K, V], key K, value V) bool {
	for k, v := range seq {
		if k == key && v == value {
			return true
		}
	}
	return false
}

// AnyBy returns whether seq has an element for that f returns true.
func AnyBy[T any](seq iter.Seq[T], f func(T) bool) bool {
	for v := range seq {
		if f(v) {
			return true
		}
	}
	return false
}

// AnyBy2 returns whether seq has an element for that f returns true.
func AnyBy2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) bool {
	for k, v := range seq {
		if f(k, v) {
			return true
		}
	}
	return false
}

// All returns whether all elements of seq are value.
func All[T comparable](seq iter.Seq[T], value T) bool {
	for v := range seq {
		if v != value {
			return false
		}
	}
	return true
}

// All2 returns whether all elements of seq are key-value pair.
func All2[K, V comparable](seq iter.Seq2[K, V], key K, value V) bool {
	for k, v := range seq {
		if k != key || v != value {
			return false
		}
	}
	return true
}

// AllBy returns whether f returns true for all elements in seq.
func AllBy[T any](seq iter.Seq[T], f func(T) bool) bool {
	for v := range seq {
		if !f(v) {
			return false
		}
	}
	return true
}

// AllBy2 returns whether f returns true for all elements in seq.
func AllBy2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) bool {
	for k, v := range seq {
		if !f(k, v) {
			return false
		}
	}
	return true
}

// Max returns the maximum element of seq.
// It returns [optional.None] if seq is empty.
func Max[T cmp.Ordered](seq iter.Seq[T]) optional.Optional[T] {
	var m T
	var init bool
	for v := range seq {
		if !init {
			m = v
			init = true
			continue
		}
		m = max(m, v)
	}
	if !init {
		return optional.None[T]()
	}
	return optional.New(m)
}

// Min returns the minimum element of seq.
// It returns [optional.None] if seq is empty.
func Min[T cmp.Ordered](seq iter.Seq[T]) optional.Optional[T] {
	var m T
	var init bool
	for v := range seq {
		if !init {
			m = v
			init = true
			continue
		}
		m = min(m, v)
	}
	if !init {
		return optional.None[T]()
	}
	return optional.New(m)
}
