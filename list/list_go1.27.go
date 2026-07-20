//go:build go1.27

package list

// Filter iterates over elements of collection, returning a list of all elements predicate returns true for.
func (l *List[T]) Filter(f func(index int, value T) bool) *List[T] {
	var ret List[T]
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if f(i, e.Value) {
			ret.PushBack(e.Value)
		}
	}
	return &ret
}

// FilterFalse iterates over elements of collection, returning a list of all elements predicate returns false for.
func (l *List[T]) FilterFalse(f func(index int, value T) bool) *List[T] {
	var ret List[T]
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if !f(i, e.Value) {
			ret.PushBack(e.Value)
		}
	}
	return &ret
}

// CountBy counts the number of elements that counter returns true.
func (l *List[T]) CountBy(f func(index int, value T) bool) int {
	var count int
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if f(i, e.Value) {
			count++
		}
	}
	return count
}

// Map returns a list of the results of applying f to each element of l.
func (l *List[T]) Map[U any](f func(index int, value T) U) *List[U] {
	var ret List[U]
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		v := f(i, e.Value)
		ret.PushBack(v)
	}
	return &ret
}
