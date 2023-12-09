package hi

import (
	"testing"
	"reflect"
)

func TestSliceIter(t *testing.T) {
	s := []int{1, 2, 3}
	j := 0
	for i, v := range SliceIter(s) {
		if i != j {
			t.Errorf("got %v, want %v", i, j)
		}
		j++

		if v != s[i] {
			t.Errorf("got %v, want %v", v, s[i])
		}
	}
}


func TestSliceValues(t *testing.T) {
	s := []int{1, 2, 3}
	i := 0
	for v := range SliceValues(s) {
		if v != s[i] {
			t.Errorf("got %v, want %v", v, s[i])
		}
		i++
	}
}

func TestMapMapIter(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	got := make(map[int]string)
	for k, v := range MapIter(m) {
		got[k] = v
	}
	if !reflect.DeepEqual(got, m) {
		t.Errorf("got %v, want %v", got, m)
	}
}

func TestMapKeys(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	got := make(map[int]bool)
	for k := range MapKeys(m) {
		got[k] = true
	}
	want := map[int]bool{1: true, 2: true, 3: true}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMapValues(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	got := make(map[string]bool)
	for v := range MapValues(m) {
		got[v] = true
	}
	want := map[string]bool{"a": true, "b": true, "c": true}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
