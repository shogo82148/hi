package hi

import (
	"github.com/shogo82148/hi/cmp"
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

// Max returns the maximum element of s.
// If s is empty, Max returns None.
// If T is a floating-point type and any of the arguments are NaNs, max will return NaN.
func Max[T cmp.Ordered](s ...T) optional.Optional[T] {
	if len(s) == 0 {
		return optional.None[T]()
	}
	max := s[0]
	for _, v := range s[1:] {
		if isNaN(v) {
			max = v
			break
		}
		if v > max {
			max = v
		}
	}
	return optional.New(max)
}

// MaxBy returns an element that f returns maximum value in s.
func MaxBy[T any, U cmp.Ordered](f func(element T) U, s ...T) optional.Optional[T] {
	if len(s) == 0 {
		return optional.None[T]()
	}
	max := s[0]
	maxW := f(s[0])
	for _, v := range s[1:] {
		if w := f(v); w > maxW {
			maxW = w
			max = v
		}
	}
	return optional.New(max)
}

// Min returns the minimum element in s.
// If s is empty, Max returns None.
// If T is a floating-point type and any of the arguments are NaNs, max will return NaN.
func Min[T cmp.Ordered](s ...T) optional.Optional[T] {
	if len(s) == 0 {
		return optional.None[T]()
	}
	min := s[0]
	for _, v := range s[1:] {
		if isNaN(v) {
			min = v
			break
		}
		if v < min {
			min = v
		}
	}
	return optional.New(min)
}

// MinBy returns an element that f returns minimum value in s.
func MinBy[T any, U cmp.Ordered](f func(element T) U, s ...T) optional.Optional[T] {
	if len(s) == 0 {
		return optional.None[T]()
	}
	min := s[0]
	minW := f(s[0])
	for _, v := range s[1:] {
		if w := f(v); w < minW {
			minW = w
			min = v
		}
	}
	return optional.New(min)
}

// MinMax returns the minimum element and the maximum element of s.
func MinMax[T cmp.Ordered](s ...T) (min optional.Optional[T], max optional.Optional[T]) {
	if len(s) == 0 {
		return
	}
	var i int
	var myMin, myMax T
	if len(s)%2 == 0 {
		x := s[0]
		y := s[1]
		if x > y {
			x, y = y, x
		}
		myMin = x
		myMax = y
		i = 2
	} else {
		myMin = s[0]
		myMax = s[0]
		i = 1
	}
	for ; i+1 < len(s); i += 2 {
		x := s[i+0]
		y := s[i+1]
		if x > y {
			x, y = y, x
		}
		if x < myMin {
			myMin = x
		}
		if y > myMax {
			myMax = y
		}
	}
	min = optional.New(myMin)
	max = optional.New(myMax)
	return
}

// MinMaxBy returns the minimum element and the maximum element of s.
func MinMaxBy[T any, U cmp.Ordered](f func(element T) U, s ...T) (min optional.Optional[T], max optional.Optional[T]) {
	if len(s) == 0 {
		return
	}
	var i int
	var myMin, myMax T
	var minW, maxW U
	if len(s)%2 == 0 {
		x, y := s[0], s[1]
		xw, yw := f(x), f(y)
		if xw > yw {
			x, xw, y, yw = y, yw, x, xw
		}
		myMin, minW = x, xw
		myMax, maxW = y, yw
		i = 2
	} else {
		w := f(s[0])
		myMin, minW = s[0], w
		myMax, maxW = s[0], w
		i = 1
	}
	for ; i+1 < len(s); i += 2 {
		x, y := s[i+0], s[i+1]
		xw, yw := f(x), f(y)
		if xw > yw {
			x, xw, y, yw = y, yw, x, xw
		}
		if xw < minW {
			myMin, minW = x, xw
		}
		if yw > maxW {
			myMax, maxW = y, yw
		}
	}
	min = optional.New(myMin)
	max = optional.New(myMax)
	return
}

// Sum returns a sum of s using Kahan summation algorithm.
func Sum[T Addable](s []T) optional.Optional[T] {
	if len(s) == 0 {
		return optional.None[T]()
	}

	sum := s[0]
	var c T
	for _, v := range s[1:] {
		y := v - c
		t := sum + y
		c = (t - sum) - y
		sum = t
	}
	return optional.New(sum)
}

// Sum returns a sum of s using Kahan summation algorithm.
func SumBy[T any, R Addable](s []T, f func(T) R) optional.Optional[R] {
	if len(s) == 0 {
		return optional.None[R]()
	}

	sum := f(s[0])
	var c R
	for _, v := range s[1:] {
		y := f(v) - c
		t := sum + y
		c = (t - sum) - y
		sum = t
	}
	return optional.New(sum)
}

// Reduce reduces s.
func Reduce[T any, U any](s []T, f func(i int, agg U, item T) U, init U) U {
	for i, v := range s {
		init = f(i, init, v)
	}
	return init
}

// ReduceRight reduces s.
func ReduceRight[T any, U any](s []T, f func(i int, agg U, item T) U, init U) U {
	for i := len(s); i > 0; i-- {
		init = f(i-1, init, s[i-1])
	}
	return init
}

// isNaN reports whether x is a NaN without requiring the math package.
// This will always return false if T is not floating-point.
func isNaN[T cmp.Ordered](x T) bool {
	return x != x
}
