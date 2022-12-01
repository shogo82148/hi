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
