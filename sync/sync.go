// The package sync provides same features as the standard sync package,
// but with generics.
package sync

import "sync"

// Cond is an alias of [sync.Cond].
type Cond = sync.Cond

// Locker is an alias of [sync.Locker].
type Locker = sync.Locker

// Mutex is an alias of [sync.Mutex].
type Mutex = sync.Mutex

// Once is an alias of [sync.Once].
type Once = sync.Once

// Pool is an alias of [sync.Pool].
type RWMutex = sync.RWMutex

// WaitGroup is an alias of [sync.WaitGroup].
type WaitGroup = sync.WaitGroup

// OnceFunc is an alias of [sync.OnceFunc].
func OnceFunc(f func()) func() {
	return sync.OnceFunc(f)
}

// OnceValue is an alias of [sync.OnceValue].
func OnceValue[T any](f func() T) func() T {
	return sync.OnceValue(f)
}

func OnceValues[T1, T2 any](f func() (T1, T2)) func() (T1, T2) {
	return sync.OnceValues(f)
}
