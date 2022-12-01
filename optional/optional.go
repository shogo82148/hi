package optional

// Optional is an optional of T.
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

// GetOrFunc returns the value of v if v is valid.
// Otherwise it calls f and return the result.
func (v Optional[T]) GetOrFunc(f func() T) T {
	if v.valid {
		return v.value
	} else {
		return f()
	}
}
