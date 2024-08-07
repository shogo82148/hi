package sets

import (
	"github.com/shogo82148/hi/tuple"
)

type Set[T comparable] map[T]struct{}

// New returns a new set that contains items.
func New[T comparable](items ...T) Set[T] {
	set := make(Set[T], len(items))
	for _, v := range items {
		set[v] = struct{}{}
	}
	return set
}

// Add adds items into the set.
func (set Set[T]) Add(items ...T) Set[T] {
	for _, v := range items {
		set[v] = struct{}{}
	}
	return set
}

// Remove removes items from the set.
func (set Set[T]) Remove(items ...T) Set[T] {
	for _, v := range items {
		delete(set, v)
	}
	return set
}

// Contains reports whether the set contains the item.
func (set Set[T]) Contains(item T) bool {
	_, ok := set[item]
	return ok
}

// Len returns the number of items in the set.
func (set Set[T]) Len() int {
	return len(set)
}

// Equal reports the other has same items.
func (set Set[T]) Equal(other Set[T]) bool {
	if len(set) != len(other) {
		return false
	}

	for v := range set {
		if _, ok := other[v]; !ok {
			return false
		}
	}
	return true
}

// Clone returns a new set that has same items with the set.
func (set Set[T]) Clone() Set[T] {
	ret := make(Set[T], len(set))
	for v := range set {
		ret[v] = struct{}{}
	}
	return ret
}

// Union returns a new set that with all items that are included by the set or the other.
func (set Set[T]) Union(other Set[T]) Set[T] {
	ret := make(Set[T], len(set)+len(other))
	for v := range set {
		ret[v] = struct{}{}
	}
	for v := range other {
		ret[v] = struct{}{}
	}
	return ret
}

// Intersection returns a new set with items that are included by both of the set and the other.
func (set Set[T]) Intersection(other Set[T]) Set[T] {
	a, b := set, other
	if len(a) > len(b) {
		// Minimize the number of the loops.
		a, b = b, a
	}

	ret := make(Set[T], len(a))
	for v := range a {
		if _, ok := b[v]; ok {
			ret[v] = struct{}{}
		}
	}
	return ret
}

// Difference returns a new set with items that are included by the set but not the other.
func (set Set[T]) Difference(other Set[T]) Set[T] {
	ret := make(Set[T], len(set))
	for v := range set {
		if _, ok := other[v]; !ok {
			ret[v] = struct{}{}
		}
	}
	return ret
}

// SymmetricDifference returns a new set.
func (set Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	ret := make(Set[T], len(set)+len(other))
	for v := range set {
		if _, ok := other[v]; !ok {
			ret[v] = struct{}{}
		}
	}
	for v := range other {
		if _, ok := set[v]; !ok {
			ret[v] = struct{}{}
		}
	}
	return ret
}

// IsSubset reports whether the set is a subset of the other.
func (set Set[T]) IsSubset(other Set[T]) bool {
	if len(set) > len(other) {
		return false
	}
	for v := range set {
		if _, ok := other[v]; !ok {
			return false
		}
	}
	return true
}

// IsProperSubset reports whether the set is a proper subset of the other.
func (set Set[T]) IsProperSubset(other Set[T]) bool {
	return len(set) < len(other) && set.IsSubset(other)
}

// IsSuperset reports whether the set is a superset of the other.
func (set Set[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(set)
}

// IsProperSuperset reports whether the set is a proper superset of the other.
func (set Set[T]) IsProperSuperset(other Set[T]) bool {
	return len(set) > len(other) && other.IsSubset(set)
}

// For calls the f for each items in the set.
// If the f returns an error, For stops the iteration and return the error.
func (set Set[T]) For(f func(v T) error) error {
	for v := range set {
		if err := f(v); err != nil {
			return err
		}
	}
	return nil
}

// Filter returns a new subset where f returns true for.
func Filter[T comparable](set Set[T], f func(v T) bool) Set[T] {
	ret := make(Set[T], len(set))
	for v := range set {
		if f(v) {
			ret[v] = struct{}{}
		}
	}
	return ret
}

// FilterFalse returns a new subset where f returns false for.
func FilterFalse[T comparable](set Set[T], f func(v T) bool) Set[T] {
	ret := make(Set[T], len(set))
	for v := range set {
		if !f(v) {
			ret[v] = struct{}{}
		}
	}
	return ret
}

// Count counts the number of elements in set that compare equal to value.
func Count[T comparable](set Set[T], value T) int {
	if set.Contains(value) {
		return 1
	}
	return 0
}

// CountBy counts the number of elements that counter returns true.
func CountBy[T comparable](set Set[T], f func(v T) bool) int {
	var count int
	for v := range set {
		if f(v) {
			count++
		}
	}
	return count
}

// Any returns whether set has value.
func Any[T comparable](set Set[T], value T) bool {
	return set.Contains(value)
}

// AnyBy returns whether set has an element for that f returns true.
func AnyBy[T comparable](set Set[T], f func(v T) bool) bool {
	for v := range set {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns whether all elements in set are value.
func All[T comparable](set Set[T], value T) bool {
	return len(set) == 0 || len(set) == 1 && set.Contains(value)
}

// AllBy returns whether f returns true for all elements in set.
func AllBy[T comparable](set Set[T], f func(v T) bool) bool {
	for v := range set {
		if !f(v) {
			return false
		}
	}
	return true
}

// Map converts all items in the set by using the mapper.
func Map[T, U comparable](set Set[T], mapper func(v T) U) Set[U] {
	ret := make(Set[U], len(set))
	for v := range set {
		ret[mapper(v)] = struct{}{}
	}
	return ret
}

// Product returns a new set.
func Product[T1, T2 comparable](s1 Set[T1], s2 Set[T2]) Set[tuple.Tuple2[T1, T2]] {
	ret := make(Set[tuple.Tuple2[T1, T2]], len(s1)*len(s2))
	for v1 := range s1 {
		for v2 := range s2 {
			ret[tuple.Tuple2[T1, T2]{V1: v1, V2: v2}] = struct{}{}
		}
	}
	return ret
}
