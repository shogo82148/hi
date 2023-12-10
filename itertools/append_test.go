//go:build goexperiment.rangefunc

package itertools

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			yield(i)
		}
	}
	got := Append([]int{}, seq)
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
