//go:build goexperiment.rangefunc

package hi

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
