package result

import "errors"

//go:generate ./generate-zip.pl
//go:generate ./generate-unzip.pl

// ErrFiltered is returned when the result is filtered.
var ErrFiltered = errors.New("result: filtered")

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
	if v.err == nil {
		return v.value
	} else {
		panic(v.err)
	}
}

func (v Result[T]) Err() error {
	return v.err
}

type wrapErrors struct {
	msg  string
	errs []error
}

func (e *wrapErrors) Error() string {
	return e.msg
}

func (e *wrapErrors) Unwrap() []error {
	return e.errs
}

func (v Result[T]) Filter(f func(T) bool) Result[T] {
	if v.err != nil {
		return v
	}
	if f(v.value) {
		return v
	}
	return Result[T]{
		err: ErrFiltered,
	}
}

func (v Result[T]) FilterFalse(f func(T) bool) Result[T] {
	if v.err != nil {
		return v
	}
	if !f(v.value) {
		return v
	}
	return Result[T]{
		err: ErrFiltered,
	}
}

func Map[T, U any](r Result[T], f func(T) U) Result[U] {
	if r.err != nil {
		return Result[U]{err: r.err}
	}
	return Result[U]{value: f(r.value)}
}
