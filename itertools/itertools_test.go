//go:build goexperiment.rangefunc

package itertools

import (
	"iter"
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

func TestAccumulate(t *testing.T) {
	seq := hi.SliceValues([]int{1, 2, 3, 4})
	got := make([]int, 0, 4)
	add := func(a, b int) int { return a + b }
	for v := range Accumulate(seq, add) {
		got = append(got, v)
	}
	want := []int{1, 3, 6, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestChain(t *testing.T) {
	seq := Chain(
		hi.SliceValues([]int{1, 2, 3, 4}),
		hi.SliceValues([]int{5, 6, 7, 8}),
	)
	got := Append(make([]int, 0, 8), seq)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestChainFromIterables(t *testing.T) {
	seq := ChainFromIterables(hi.SliceValues([]iter.Seq[int]{
		hi.SliceValues([]int{1, 2, 3, 4}),
		hi.SliceValues([]int{5, 6, 7, 8}),
	}))
	got := Append(make([]int, 0, 8), seq)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCompress(t *testing.T) {
	seq := Compress(
		hi.SliceValues([]int{1, 2, 3, 4}),
		hi.SliceValues([]bool{true, false, true, false}),
	)
	got := Append(make([]int, 0, 2), seq)
	want := []int{1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDropWhile(t *testing.T) {
	seq := DropWhile(
		func(v int) bool { return v < 5 },
		hi.SliceValues([]int{1, 4, 6, 4, 1}),
	)
	got := Append(make([]int, 0, 2), seq)
	want := []int{6, 4, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFilterFalse(t *testing.T) {
	seq := FilterFalse(
		func(v int) bool { return v % 2 != 0 },
		Range(10),
	)
	got := Append(make([]int, 0, 5), seq)
	want := []int{0, 2, 4, 6, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
