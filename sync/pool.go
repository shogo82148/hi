package sync

import "sync"

// A Pool is a generics-enabled version of [sync.Pool].
type Pool[T any] struct {
	// New optionally specifies a function to generate
	// a value when Get would otherwise return zero value.
	// It may not be changed concurrently with calls to Get.
	New func() T

	pool sync.Pool
}

// Get selects an arbitrary item from the Pool, removes it from the Pool, and returns it to the caller.
// Get may choose to ignore the pool and treat it as empty.
// Callers should not assume any relation between values passed to Put and the values returned by Get.
//
// If Get would otherwise return zero value and p.New is non-nil, Get returns the result of calling p.New.
func (p *Pool[T]) Get() T {
	v := p.pool.Get()
	if v == nil {
		if p.New != nil {
			return p.New()
		}

		var zero T
		return zero
	}
	return v.(T)
}

// Put adds x to the pool.
func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}
