//go:build goexperiment.rangefunc

package itertools

import (
	"reflect"
	"testing"

	"github.com/shogo82148/hi"
)

func TestCount(t *testing.T) {
	got := make([]int, 0, 5)
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

func TestCycle(t *testing.T) {
	got := make([]int, 0, 10)
	seq := hi.SliceValues([]int{1, 2, 3, 4})
	for v := range Cycle(seq) {
		got = append(got, v)
		if len(got) >= 10 {
			break
		}
	}
	want := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRepeat(t *testing.T) {
	seq := Repeat(123)

	got := make([]int, 0, 5)
	for v := range seq {
		got = append(got, v)
		if len(got) >= 5 {
			break
		}
	}
	want := []int{123, 123, 123, 123, 123}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRepeatN(t *testing.T) {
	seq := RepeatN(123, 5)

	got := Append(make([]int, 0, 5), seq)
	want := []int{123, 123, 123, 123, 123}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
