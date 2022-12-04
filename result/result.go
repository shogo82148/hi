package result

type Result[T any] struct {
	value T
	err   error
}

func New[T any](v T, err error) Result[T] {
	if err != nil {
		return Result[T]{
			err: err,
		}
	}
	return Result[T]{
		value: v,
	}
}

func OK[T any](v T) Result[T] {
	return Result[T]{
		value: v,
	}
}

func Error[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

func (v Result[T]) Get() T {
	if v.err != nil {
		return v.value
	} else {
		panic(v.err)
	}
}

func (v Result[T]) Err() error {
	return v.err
}
