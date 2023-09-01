package sync

import "testing"

// We assume that the Pool tests have already covered parallelism.

func TestPool(t *testing.T) {
	p1 := &Pool[string]{}
	p1.Get()
	p1.Put("hello")

	p2 := &Pool[string]{
		New: func() string {
			return "world"
		},
	}
	if p2.Get() != "world" {
		t.Fatal("Get returned wrong value")
	}
}
