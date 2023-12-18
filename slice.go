package hi

import (
	"cmp"
	"slices"

	"github.com/shogo82148/hi/optional"
)

//go:generate ./generate-zip.pl
//go:generate ./generate-unzip.pl

// Addable is a type that can be added and subtracted.
type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

// Map calls mapper on each element of a and returns the result of its result in a slice.
func Map[S1 ~[]T1, S2 []T2, T1, T2 any](a S1, mapper func(int, T1) T2) S2 {
	ret := make(S2, len(a))
	for i, v := range a {
		ret[i] = mapper(i, v)
	}
	return ret
}

// Filter iterates over elements of collection, returning a slice of all elements predicate returns true for.
func Filter[S ~[]T, T any](a S, filter func(int, T) bool) S {
	ret := make(S, 0, len(a))
	for i, v := range a {
		if filter(i, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// FilterFalse iterates over elements of collection, returning a slice of all elements predicate returns false for.
func FilterFalse[S ~[]T, T any](a S, filter func(int, T) bool) S {
	ret := make(S, 0, len(a))
	for i, v := range a {
		if !filter(i, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// GroupBy returns a map of slices, where each slice contains elements of a grouped by the result of f.
func GroupBy[S ~[]T, T any, K comparable](a S, f func(int, T) K) map[K]S {
	ret := make(map[K]S)
	for i, v := range a {
		k := f(i, v)
		ret[k] = append(ret[k], v)
	}
	return ret
}

// Chunk creates a slice of elements split into groups the length of size.
func Chunk[S ~[]T, T any](a S, size int) []S {
	if size <= 0 {
		panic("size must be positive")
	}
	ret := make([]S, 0, (len(a)+size-1)/size)
	for i := 0; i < len(a); i += size {
		end := i + size
		if end > len(a) {
			end = len(a)
		}
		ret = append(ret, a[i:end])
	}
	return ret
}

// Slice returns a Python-like slice of s.
// Unlike Go's builtin slice, Slice can handle negative indices and step,
// and always returns a new slice.
// Negative indices count from the end of the slice (-1 means the last element).
func Slice[S ~[]T, T any](s S, start, stop, step int) S {
	if start < 0 {
		start += len(s)
	}
	if start < 0 {
		start = 0
	} else if start > len(s) {
		start = len(s)
	}

	if stop < 0 {
		stop += len(s)
	}
	if stop < 0 {
		stop = 0
	} else if stop > len(s) {
		stop = len(s)
	}

	if step == 0 {
		panic("step must not be zero")
	}

	if step < 0 {
		ret := make(S, 0, (start-stop-step-1)/(-step))
		for i := start; i > stop; i += step {
			ret = append(ret, s[i])
		}
		return ret
	} else {
		ret := make(S, 0, (stop-start+step-1)/step)
		for i := start; i < stop; i += step {
			ret = append(ret, s[i])
		}
		return ret
	}
}

// Chain returns a slice of all elements from all slices.
func Chain[S ~[]T, T any](a ...S) []T {
	var ret []T
	for _, v := range a {
		ret = append(ret, v...)
	}
	return ret
}

// Compress makes a slice of the elements in a for which the corresponding element in selectors is true.
func Compress[S ~[]T, T any](a S, selectors []bool) S {
	l := min(len(a), len(selectors))
	ret := make(S, 0, l)
	for i := 0; i < l; i++ {
		if selectors[i] {
			ret = append(ret, a[i])
		}
	}
	return ret
}

// DropWhile returns a slice of the remaining elements after dropping the longest prefix of elements that satisfy predicate.
func DropWhile[S ~[]T, T any](a S, predicate func(int, T) bool) S {
	for i, v := range a {
		if !predicate(i, v) {
			return slices.Clone(a[i:])
		}
	}
	return make(S, 0)
}

// TakeWhile returns a slice of the longest prefix of elements that satisfy predicate f.
func TakeWhile[S ~[]T, T any](a S, f func(int, T) bool) S {
	for i, v := range a {
		if !f(i, v) {
			return slices.Clone(a[:i])
		}
	}
	return slices.Clone(a)
}

// RepeatN returns a slice consisting of n copies of v.
func RepeatN[T any](v T, n int) []T {
	if n < 0 {
		panic("n must be positive")
	}
	ret := make([]T, n)
	for i := range ret {
		ret[i] = v
	}
	return ret
}

// Count counts the number of elements in the collection that compare equal to value.
func Count[T comparable](a []T, value T) int {
	var count int
	for _, v := range a {
		if v == value {
			count++
		}
	}
	return count
}

// CountBy counts the number of elements that counter returns true.
func CountBy[T any](a []T, counter func(int, T) bool) int {
	var count int
	for i, v := range a {
		if counter(i, v) {
			count++
		}
	}
	return count
}

// Any returns whether a has value at least one.
func Any[T comparable](a []T, value T) bool {
	for _, v := range a {
		if v == value {
			return true
		}
	}
	return false
}

// AnyBy returns whether a has an element for that f returns true.
func AnyBy[T any](a []T, f func(index int, value T) bool) bool {
	for i, v := range a {
		if f(i, v) {
			return true
		}
	}
	return false
}

// All returns whether all elements of a are value.
func All[T comparable](a []T, value T) bool {
	for _, v := range a {
		if v != value {
			return false
		}
	}
	return true
}

// AllBy returns whether f returns true for all elements in a.
func AllBy[T any](a []T, f func(int, T) bool) bool {
	for i, v := range a {
		if !f(i, v) {
			return false
		}
	}
	return true
}

// Max returns the maximal value in a.
func Max[S ~[]T, T cmp.Ordered](a S) optional.Optional[T] {
	if len(a) == 0 {
		return optional.None[T]()
	}
	return optional.New(slices.Max(a))
}

// Min returns the minimal value in a.
func Min[S ~[]T, T cmp.Ordered](a S) optional.Optional[T] {
	if len(a) == 0 {
		return optional.None[T]()
	}
	return optional.New(slices.Min(a))
}
