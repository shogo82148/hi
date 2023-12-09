package iter

// Slice returns a slice of the elements of the iterator.
func Slice[S ~[]E, E any](iter func (func(E)bool)) S {
	s := make(S, 0)
	for v := range iter {
		s = append(s, v)
	}
	return s
}
