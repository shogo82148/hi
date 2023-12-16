package hi

func Negate[F ~func(T)bool, T any](f F) F {
	return func(v T) bool {
		return !f(v)
	}
}

func Negate2[F ~func(T, U)bool, T, U any](f F) F {
	return func(v T, w U) bool {
		return !f(v, w)
	}
}
