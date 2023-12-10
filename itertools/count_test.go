package itertools

import (
	"testing"
	"reflect"
)

func TestCount(t *testing.T) {
	got := make([]int, 0)
	for v := range Count(0, 2) {
		got = append(got, v)
		if len(got) >= 5 {
			break
		}
	}
	want := []int{0, 2, 4, 6, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
