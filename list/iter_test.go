package list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestForward(t *testing.T) {
	l := New[string]()
	for i := 0; i < 5; i++ {
		l.PushBack(fmt.Sprintf("%d", i))
	}

	got := make([]string, 0, l.Len())
	for _, v := range l.Forward() {
		got = append(got, v)
	}
	want := []string{"0", "1", "2", "3", "4"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBackward(t *testing.T) {
	l := New[string]()
	for i := 0; i < 5; i++ {
		l.PushBack(fmt.Sprintf("%d", i))
	}

	got := make([]string, 0, l.Len())
	for _, v := range l.Backward() {
		got = append(got, v)
	}
	want := []string{"4", "3", "2", "1", "0"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
