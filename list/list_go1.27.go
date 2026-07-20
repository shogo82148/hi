//go:build go1.27

package list

// Filter returns the elements of l for which f returns true.
func (l *List[T]) Filter(f func(index int, value T) bool) *List[T] {
	var ret List[T]
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if f(i, e.Value) {
			ret.PushBack(e.Value)
		}
	}
	return &ret
}

// FilterFalse returns the elements of l for which f returns false.
func (l *List[T]) FilterFalse(f func(index int, value T) bool) *List[T] {
	var ret List[T]
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if !f(i, e.Value) {
			ret.PushBack(e.Value)
		}
	}
	return &ret
}

// CountBy returns the number of elements of l for which f returns true.
func (l *List[T]) CountBy(f func(index int, value T) bool) int {
	var count int
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if f(i, e.Value) {
			count++
		}
	}
	return count
}

// AnyBy reports whether any element of l makes f return true.
func (l *List[T]) AnyBy(f func(index int, value T) bool) bool {
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if f(i, e.Value) {
			return true
		}
	}
	return false
}

// AllBy reports whether all elements of l make f return true.
func (l *List[T]) AllBy(f func(int, T) bool) bool {
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		if !f(i, e.Value) {
			return false
		}
	}
	return true
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

// Reduce reduces l to a single value which is the accumulated result of running each element in l through the reducer function.
func (l *List[T]) Reduce[U any](f func(i int, acc U, item T) U, init U) U {
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		init = f(i, init, e.Value)
	}
	return init
}
