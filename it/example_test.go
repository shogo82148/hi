//go:build goexperiment.rangefunc

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

func ExampleZip() {
	seq := it.Zip(it.Range(5), it.SliceValues([]string{"a", "b", "c"}))
	for k, v := range seq {
		fmt.Printf("(%d, %s)\n", k, v)
	}

	// Output:
	// (0, a)
	// (1, b)
	// (2, c)
}

func ExampleZip2() {
	seq := it.Zip2(it.Range(5), it.SliceValues([]string{"a", "b", "c"}))
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// (0, a)
	// (1, b)
	// (2, c)
}

func ExampleMap() {
	seq := it.Map(it.Range(5), func(v int) string {
		return fmt.Sprintf("(%d)", v)
	})
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// (0)
	// (1)
	// (2)
	// (3)
	// (4)
}

func ExampleMap2() {
	in := it.SliceIter([]string{"a", "b", "c"})
	seq := it.Map2(in, func(i int, v string) (int, string) {
		return i, fmt.Sprintf("(%s)", v)
	})
	for i, v := range seq {
		fmt.Printf("%d: %s\n", i, v)
	}

	// Output:
	// 0: (a)
	// 1: (b)
	// 2: (c)
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

func ExampleGroupBy() {
	seq := it.GroupBy(it.Range(5), func(v int) int {
		return v % 2
	})
	for k, v := range seq {
		fmt.Printf("%d: %v\n", k, v)
	}
	// Unordered output:
	// 0: [0 2 4]
	// 1: [1 3]
}

func ExampleChunk() {
	seq := it.Chunk(it.Range(5), 2)
	for v := range seq {
		fmt.Println(v)
	}
	// Output:
	// [0 1]
	// [2 3]
	// [4]
}

func ExampleCount() {
	fmt.Println(it.Count(it.Range(5), 3))
	// Output:
	// 1
}

func ExampleCountBy() {
	fmt.Println(it.CountBy(it.Range(5), func(v int) bool { return v%2 == 0 }))
	// Output:
	// 3
}
