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
