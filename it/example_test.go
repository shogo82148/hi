package it_test

import (
	"fmt"

	"github.com/shogo82148/hi/it"
)

func ExampleSliceIter() {
	seq := it.SliceIter([]string{"a", "b", "c"})
	for i, v := range seq {
		fmt.Printf("%d: %s\n", i, v)
	}

	// Output:
	// 0: a
	// 1: b
	// 2: c
}

func ExampleSliceValues() {
	seq := it.SliceValues([]string{"a", "b", "c"})
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// a
	// b
	// c
}

func ExampleMapIter() {
	seq := it.MapIter(map[string]int{"a": 1, "b": 2, "c": 3})
	for k, v := range seq {
		fmt.Printf("%s: %d\n", k, v)
	}

	// Unordered output:
	// a: 1
	// b: 2
	// c: 3
}

func ExampleMapKeys() {
	seq := it.MapKeys(map[string]int{"a": 1, "b": 2, "c": 3})
	for v := range seq {
		fmt.Println(v)
	}

	// Unordered output:
	// a
	// b
	// c
}

func ExampleMapValues() {
	seq := it.MapValues(map[string]int{"a": 1, "b": 2, "c": 3})
	for v := range seq {
		fmt.Println(v)
	}
	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleChanValues() {
	ch := make(chan string, 3)
	ch <- "a"
	ch <- "b"
	ch <- "c"
	close(ch)

	seq := it.ChanValues(ch)
	for v := range seq {
		fmt.Println(v)
	}
	// Unordered output:
	// a
	// b
	// c
}

func ExampleFilter() {
	seq := it.Filter(it.Range(5), func(v int) bool {
		return v%2 == 0
	})
	for v := range seq {
		fmt.Println(v)
	}
	// Output:
	// 0
	// 2
	// 4
}
