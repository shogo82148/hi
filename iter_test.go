package hi

import "testing"

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
