package result

import (
	"errors"
	"testing"
)

func TestGet(t *testing.T) {
	var ret Result[int]

	// OK
	ret = OK(1)
	if ret.Get() != 1 {
		t.Errorf("got %d, want %d", ret.Get(), 1)
	}

	// Error
	ret = Error[int](errors.New("some error"))
	var panicked bool
	func() {
		defer func() {
			if err := recover(); err != nil {
				panicked = true
			}
		}()
		ret.Get()
	}()
	if !panicked {
		t.Error("want to panic, but not")
	}
}
