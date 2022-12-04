package optional

//go:generate ./generate-zip.pl
//go:generate ./generate-unzip.pl

// Optional is an optional of T.
// Zero value equals None[T]()
type Optional[T any] struct {
	value T
	valid bool
}

// New return a new value.
func New[T any](v T) Optional[T] {
	return Optional[T]{
		value: v,
		valid: true,
	}
}

// None returns nothing.
func None[T any]() Optional[T] {
	return Optional[T]{}
}

// Valid reports v is valid.
func (v Optional[T]) Valid() bool {
	return v.valid
}

// Get returns the value of v if v is valid.
// Otherwise it panics.
func (v Optional[T]) Get() T {
	if v.valid {
		return v.value
	} else {
		panic("get a value from none")
	}
}

// Get returns the value of v if v is valid.
// Otherwise it returns defaults.
func (v Optional[T]) GetOr(defaults T) T {
	if v.valid {
		return v.value
	} else {
		return defaults
	}
}

// Get returns the value of v if v is valid.
// Otherwise it returns zero value.
func (v Optional[T]) GetOrZero() T {
	if v.valid {
		return v.value
	} else {
		var zero T
		return zero
	}
}

// GetOrFunc returns the value of v if v is valid.
// Otherwise it calls f and returns its result.
func (v Optional[T]) GetOrFunc(f func() T) T {
	if v.valid {
		return v.value
	} else {
		return f()
	}
}
