package it

import (
	"iter"
	"reflect"
	"testing"
)

func TestCycle(t *testing.T) {
	got := make([]int, 0, 10)
	seq := SliceValues([]int{1, 2, 3, 4})
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
	seq := SliceValues([]int{1, 2, 3, 4})
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
		SliceValues([]int{1, 2, 3, 4}),
		SliceValues([]int{5, 6, 7, 8}),
	)
	got := Append(make([]int, 0, 8), seq)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestChainFromIterables(t *testing.T) {
	seq := ChainFromIterables(SliceValues([]iter.Seq[int]{
		SliceValues([]int{1, 2, 3, 4}),
		SliceValues([]int{5, 6, 7, 8}),
	}))
	got := Append(make([]int, 0, 8), seq)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCompress(t *testing.T) {
	seq := Compress(
		SliceValues([]int{1, 2, 3, 4}),
		SliceValues([]bool{true, false, true, false}),
	)
	got := Append(make([]int, 0, 2), seq)
	want := []int{1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFilterFalse(t *testing.T) {
	seq := FilterFalse(
		Range(10),
		func(v int) bool { return v%2 != 0 },
	)
	got := Append(make([]int, 0, 5), seq)
	want := []int{0, 2, 4, 6, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTee(t *testing.T) {
	s := Tee(Range(3), 2)
	if len(s) != 2 {
		t.Fatalf("got %d, want %d", len(s), 2)
	}

	next0, stop0 := iter.Pull(s[0])
	defer stop0()
	next1, stop1 := iter.Pull(s[1])
	defer stop1()

	i := 0
	for {
		v0, ok0 := next0()
		v1, ok1 := next1()
		if !ok0 || !ok1 {
			break
		}
		if v0 != v1 {
			t.Errorf("got %d, want %d", v0, v1)
		}
		i++
	}
	if i != 3 {
		t.Errorf("got %d, want %d", i, 3)
	}
}

func TestTee2(t *testing.T) {
	s := Tee2(SliceIter([]string{"zero", "one", "two"}), 2)
	if len(s) != 2 {
		t.Fatalf("got %d, want %d", len(s), 2)
	}

	next0, stop0 := iter.Pull2(s[0])
	defer stop0()
	next1, stop1 := iter.Pull2(s[1])
	defer stop1()

	i := 0
	for {
		k0, v0, ok0 := next0()
		k1, v1, ok1 := next1()
		if !ok0 || !ok1 {
			break
		}
		if k0 != k1 || v0 != v1 {
			t.Errorf("got (%d, %s), want (%d, %s)", k0, v0, k1, v1)
		}
		i++
	}
	if i != 3 {
		t.Errorf("got %d, want %d", i, 3)
	}
}

func BenchmarkTee(b *testing.B) {
	s := Tee(Range(b.N), 2)
	next0, stop0 := iter.Pull(s[0])
	defer stop0()
	next1, stop1 := iter.Pull(s[1])
	defer stop1()

	i := 0
	for {
		v0, ok0 := next0()
		v1, ok1 := next1()
		if !ok0 || !ok1 {
			break
		}
		if v0 != v1 {
			b.Errorf("got %d, want %d", v0, v1)
		}
		i++
	}
	if i != b.N {
		b.Errorf("got %d, want %d", i, b.N)
	}
}

func BenchmarkTee2(b *testing.B) {
	s := Tee2(Enumerate(Range(b.N)), 2)
	next0, stop0 := iter.Pull2(s[0])
	defer stop0()
	next1, stop1 := iter.Pull2(s[1])
	defer stop1()

	i := 0
	for {
		k0, v0, ok0 := next0()
		k1, v1, ok1 := next1()
		if !ok0 || !ok1 {
			break
		}
		if k0 != k1 || v0 != v1 {
			b.Errorf("got (%d, %d), want (%d, %d)", k0, v0, k1, v1)
		}
		i++
	}
	if i != b.N {
		b.Errorf("got %d, want %d", i, b.N)
	}
}
