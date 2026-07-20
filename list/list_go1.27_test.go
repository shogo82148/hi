//go:build go1.27

package list

import (
	"fmt"
	"slices"
	"testing"
)

func TestList_Filter(t *testing.T) {
	l := New[int]()
	for i := range 5 {
		l.PushBack(i)
	}

	filtered := l.Filter(func(index int, value int) bool {
		return value%2 == 0
	})

	got := make([]int, 0, filtered.Len())
	for _, v := range filtered.Forward() {
		got = append(got, v)
	}
	want := []int{0, 2, 4}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestList_Map(t *testing.T) {
	l := New[int]()
	for i := range 5 {
		l.PushBack(i)
	}

	mapped := l.Map(func(index int, value int) string {
		return fmt.Sprintf("%d:%d", index, value)
	})

	got := make([]string, 0, mapped.Len())
	for _, v := range mapped.Forward() {
		got = append(got, v)
	}
	want := []string{"0:0", "1:1", "2:2", "3:3", "4:4"}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
