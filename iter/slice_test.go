package iter

import (
	"testing"
	"reflect"
)

func TestSlice(t *testing.T) {
	iter := func (yield func(int)bool) {
		for i := 0; i < 10; i++ {
			yield(i)
		}
	}
	got := Slice[[]int](iter)
	want := []int{0,1,2,3,4,5,6,7,8,9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
