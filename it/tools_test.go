package it

import (
	"iter"
	"math"
	"reflect"
	"slices"
	"testing"
)

func TestAppend(t *testing.T) {
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i) {
				break
			}
		}
	}
	got := Append([]int{}, seq)
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestKeys(t *testing.T) {
	seq := MapIter(map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4"})

	got := Append([]int{}, Keys(seq))
	slices.Sort(got)
	want := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestValues(t *testing.T) {
	seq := MapIter(map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4"})

	got := Append([]string{}, Values(seq))
	slices.Sort(got)
	want := []string{"0", "1", "2", "3", "4"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	seq := Filter(
		Range(10),
		func(v int) bool { return v%2 != 0 },
	)
	got := Append(make([]int, 0, 5), seq)
	want := []int{1, 3, 5, 7, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFilter2(t *testing.T) {
	seq := Filter2(
		SliceIter([]string{"0", "1", "2", "3", "4"}),
		func(k int, v string) bool { return k%2 != 0 },
	)
	got := make([]string, 0, 5)
	for _, v := range seq {
		got = append(got, v)
	}
	want := []string{"1", "3"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestUnzip(t *testing.T) {
	keys, values := Unzip(SliceIter([]string{"0", "1", "2", "3", "4"}))

	nextKey, stopKey := iter.Pull(keys)
	defer stopKey()
	nextValue, stopValue := iter.Pull(values)
	defer stopValue()

	keysGot := make([]int, 0, 5)
	valuesGot := make([]string, 0, 5)
	for {
		k, ok := nextKey()
		if !ok {
			break
		}
		v, ok := nextValue()
		if !ok {
			break
		}

		keysGot = append(keysGot, k)
		valuesGot = append(valuesGot, v)
	}

	keysWant := []int{0, 1, 2, 3, 4}
	valuesWant := []string{"0", "1", "2", "3", "4"}
	if !reflect.DeepEqual(keysGot, keysWant) {
		t.Errorf("got %v, want %v", keysGot, keysWant)
	}
	if !reflect.DeepEqual(valuesGot, valuesWant) {
		t.Errorf("got %v, want %v", valuesGot, valuesWant)
	}
}

func TestMax_Nan(t *testing.T) {
	input := []float64{
		0, 1, 2, math.NaN(), 4, 5,
	}
	result := Max(SliceValues(input))
	if !result.Valid() {
		t.Errorf("Max(%v) = %t, want %t", input, result.Valid(), true)
	}
	got := result.Get()
	if !math.IsNaN(got) {
		t.Errorf("Max(%v) = %f, want NaN", input, got)
	}
}

func TestSample(t *testing.T) {
	seq := SliceValues([]int{1, 2, 3, 4, 5})
	count := make([]int, 6)
	for i := 0; i < 100; i++ {
		ret := Sample(seq)
		if !ret.Valid() {
			t.Errorf("Sample() = %t, want %t", ret.Valid(), true)
		}
		got := ret.Get()
		if got < 1 || got > 5 {
			t.Errorf("Sample() = %d, want 1 <= x <= 5", got)
		}
		count[got]++
	}
	for i := 1; i <= 5; i++ {
		t.Logf("count[%d] = %d", i, count[i])
	}
}

func TestSample2(t *testing.T) {
	slice := []string{"zero", "one", "two", "three", "four"}
	seq := SliceIter(slice)
	count := make([]int, 5)
	for i := 0; i < 100; i++ {
		ret := Sample2(seq)
		if !ret.Valid() {
			t.Errorf("Sample2() = %t, want %t", ret.Valid(), true)
		}
		got := ret.Get()
		if got.V1 < 0 || got.V1 > 4 {
			t.Errorf("Sample2() = %d, want 0 <= x <= 4", got.V1)
		}
		if got.V2 != slice[got.V1] {
			t.Errorf("Sample2() = %s, want %s", got.V2, slice[got.V1])
		}
		count[got.V1]++
	}

	for i := 0; i < 5; i++ {
		t.Logf("count[%d] = %d", i, count[i])
	}
}

func TestSampleN(t *testing.T) {
	seq := SliceValues([]int{1, 2, 3, 4, 5})
	for i := 0; i < 100; i++ {
		ret := SampleN(seq, 3)
		if len(ret) != 3 {
			t.Errorf("SampleN() = %d, want 3", len(ret))
		}
		for _, v := range ret {
			if v < 1 || v > 5 {
				t.Errorf("SampleN() = %d, want 1 <= x <= 5", v)
			}
		}
	}
}

func TestSampleN2(t *testing.T) {
	slice := []string{"zero", "one", "two", "three", "four"}
	seq := SliceIter(slice)
	for i := 0; i < 100; i++ {
		ret := SampleN2(seq, 3)
		if len(ret) != 3 {
			t.Errorf("SampleN2() = %d, want 3", len(ret))
		}
		for _, v := range ret {
			if v.V1 < 0 || v.V1 > 4 {
				t.Errorf("SampleN2() = %d, want 0 <= x <= 4", v.V1)
			}
			if v.V2 != slice[v.V1] {
				t.Errorf("SampleN2() = %s, want %s", v.V2, slice[v.V1])
			}
		}
	}
}
