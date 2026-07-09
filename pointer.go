package hi

// Ptr returns a pointer to v.
//
//go:fix inline
func Ptr[T any](v T) *T {
	return new(v)
}
