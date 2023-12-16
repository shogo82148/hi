//go:build goexperiment.rangefunc

package it

import (
	"fmt"
	"iter"
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

func TestKeyValues(t *testing.T) {
	seq1 := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i) {
				break
			}
		}
	}
	seq2 := func(yield func(string) bool) {
		for i := 0; i < 10; i++ {
			if !yield(fmt.Sprintf("%d", i)) {
				break
			}
		}
	}

	got := make(map[int]string)
	for k, v := range KeyValues(seq1, seq2) {
		got[k] = v
	}
	want := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
	}
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

	for {
		k, ok := nextKey()
		if !ok {
			break
		}
		v, ok := nextValue()
		if !ok {
			break
		}

		fmt.Printf("%d: %s\n", k, v)
	}
}
