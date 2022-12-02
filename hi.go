package hi

//go:generate ./generate-zip.pl
//go:generate ./generate-unzip.pl

// Map calls mapper on each element of a and returns the result of its result in a slice.
func Map[T, U any](a []T, mapper func(int, T) U) []U {
	ret := make([]U, len(a))
	for i, v := range a {
		ret[i] = mapper(i, v)
	}
	return ret
}

// Filter iterates over elements of collection, returning a slice of all elements predicate returns true for.
func Filter[T any](a []T, filter func(int, T) bool) []T {
	ret := make([]T, 0, len(a))
	for i, v := range a {
		if filter(i, v) {
			ret = append(ret, v)
		}
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

// Count counts the number of elements that counter returns true.
func CountBy[T any](a []T, counter func(int, T) bool) int {
	var count int
	for i, v := range a {
		if counter(i, v) {
			count++
		}
	}
	return count
}

func Any[T comparable](a []T, value T) bool {
	for _, v := range a {
		if v == value {
			return true
		}
	}
	return false
}

func AnyBy[T any](a []T, f func(int, T) bool) bool {
	for i, v := range a {
		if f(i, v) {
			return true
		}
	}
	return false
}
