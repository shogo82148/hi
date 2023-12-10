//go:build goexperiment.rangefunc

package itertools

import (
	"iter"
)

// Append appends all elements of seq to s and returns the resulting slice.
func Append[S ~[]E, E any](s S, seq iter.Seq[E]) S {
	for v := range seq {
		s = append(s, v)
	}
	return s
}
