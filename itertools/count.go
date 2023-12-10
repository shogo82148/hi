package itertools

type Countable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64
}

// Count return an iterator that returns evenly spaced values starting with number start.
func Count[T Countable](start, step T) func(func(T) bool) {
	return func(yield func(T) bool) {
		for i := start;; i += step {
			if !yield(i) {
				break
			}
		}
	}
}
