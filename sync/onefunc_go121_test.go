//go:build go1.21
// +build go1.21

package sync_test

import (
	"runtime"
	"testing"

	"github.com/shogo82148/hi/sync"
)

func TestOnceFuncPanicNil(t *testing.T) {
	calls := 0
	f := sync.OnceFunc(func() {
		calls++
		panic(nil)
	})
	testOncePanicWith(t, &calls, f, func(label string, p any) {
		switch p.(type) {
		case nil, *runtime.PanicNilError:
			return
		}
		t.Fatalf("%s: want nil panic, got %v", label, p)
	})
}
