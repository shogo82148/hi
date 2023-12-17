package hi

// Negate returns a function that negates the result of f.
func Negate[F ~func(T) bool, T any](f F) F {
	return func(v T) bool {
		return !f(v)
	}
}

// Negate2 returns a function that negates the result of f.
func Negate2[F ~func(T, U) bool, T, U any](f F) F {
	return func(v T, w U) bool {
		return !f(v, w)
	}
}
