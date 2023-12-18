//go:build goexperiment.rangefunc

package it_test

import (
	"fmt"
	"iter"

	"github.com/shogo82148/hi/it"
	"github.com/shogo82148/hi/tuple"
)

func ExampleCycle() {
	seq := it.Cycle(it.Range(3))

	i := 0
	for v := range seq {
		i++
		if i > 10 {
			break
		}
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// 0
	// 1
	// 2
	// 0
	// 1
	// 2
	// 0
}

func ExampleCycle2() {
	seq := it.Cycle2(it.SliceIter([]string{"zero", "one", "two"}))

	i := 0
	for k, v := range seq {
		i++
		if i > 10 {
			break
		}
		fmt.Printf("%d: %s\n", k, v)
	}

	// Output:
	// 0: zero
	// 1: one
	// 2: two
	// 0: zero
	// 1: one
	// 2: two
	// 0: zero
	// 1: one
	// 2: two
	// 0: zero
}

func ExampleRange() {
	seq := it.Range(5)
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleSlice_StartEnd() {
	seq := it.SliceValues([]string{"a", "b", "c", "d", "e", "f", "g"})
	seq = it.Slice(seq, 2, 4, 1)
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// c
	// d
}

func ExampleSlice_Step() {
	seq := it.SliceValues([]string{"a", "b", "c", "d", "e", "f", "g"})
	seq = it.Slice(seq, 0, -1, 2)
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// a
	// c
	// e
	// g
}

func ExampleSlice2_StartEnd() {
	seq := it.SliceIter([]string{"a", "b", "c", "d", "e", "f", "g"})
	seq = it.Slice2(seq, 2, 4, 1)
	for i, v := range seq {
		fmt.Printf("%d: %s\n", i, v)
	}

	// Output:
	// 2: c
	// 3: d
}

func ExampleSlice2_Step() {
	seq := it.SliceIter([]string{"a", "b", "c", "d", "e", "f", "g"})
	seq = it.Slice2(seq, 0, -1, 2)
	for i, v := range seq {
		fmt.Printf("%d: %s\n", i, v)
	}

	// Output:
	// 0: a
	// 2: c
	// 4: e
	// 6: g
}

func ExampleChain() {
	seq := it.Chain(it.Range(3), it.Range(3))
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// 0
	// 1
	// 2
}

func ExampleChainFromIterables() {
	seq := func(yield func(seq iter.Seq[int]) bool) {
		if !yield(it.Range(3)) {
			return
		}
		if !yield(it.Range(3)) {
			return
		}
	}
	for v := range it.ChainFromIterables(seq) {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// 0
	// 1
	// 2
}

func ExampleChain2() {
	seq := it.Chain2(
		it.SliceIter([]string{"zero", "one", "two"}),
		it.SliceIter([]string{"null", "eins", "zwei"}),
	)
	for i, v := range seq {
		fmt.Printf("%d: %s\n", i, v)
	}

	// Output:
	// 0: zero
	// 1: one
	// 2: two
	// 0: null
	// 1: eins
	// 2: zwei
}

func ExampleChainFromIterables2() {
	seq := func(yield func(seq iter.Seq2[int, string]) bool) {
		if !yield(it.SliceIter([]string{"zero", "one", "two"})) {
			return
		}
		if !yield(it.SliceIter([]string{"null", "eins", "zwei"})) {
			return
		}
	}
	for i, v := range it.ChainFromIterables2(seq) {
		fmt.Printf("%d: %s\n", i, v)
	}

	// Output:
	// 0: zero
	// 1: one
	// 2: two
	// 0: null
	// 1: eins
	// 2: zwei
}

func ExampleCompress() {
	seq := it.Compress(it.Range(5), it.SliceValues([]bool{true, false, true, false, true}))
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 2
	// 4
}

func ExampleCompress2() {
	seq := it.Compress2(it.SliceIter([]string{"zero", "one", "two"}), it.SliceValues([]bool{true, false, true}))
	for k, v := range seq {
		fmt.Printf("%d: %s\n", k, v)
	}

	// Output:
	// 0: zero
	// 2: two
}

func ExampleDropWhile() {
	seq := it.DropWhile(
		it.SliceValues([]int{1, 4, 6, 4, 1}),
		func(v int) bool { return v < 5 },
	)
	for v := range seq {
		fmt.Println(v)
	}

	// Output:
	// 6
	// 4
	// 1
}

func ExampleDropWhile2() {
	seq := it.DropWhile2(
		it.SliceIter([]int{1, 4, 6, 4, 1}),
		func(_, v int) bool { return v < 5 },
	)
	for k, v := range seq {
		fmt.Printf("%d: %d\n", k, v)
	}

	// Output:
	// 2: 6
	// 3: 4
	// 4: 1
}

func ExampleTee() {
	seq := it.Tee(it.Range(3), 2)
	for v := range seq[0] {
		fmt.Println(v)
	}
	for v := range seq[1] {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// 0
	// 1
	// 2
}

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

func ExampleEnumerate() {
	seq := it.Enumerate(it.SliceValues([]string{"a", "b", "c"}))
	for k, v := range seq {
		fmt.Printf("%d: %s\n", k, v)
	}

	// Output:
	// 0: a
	// 1: b
	// 2: c
}

func ExampleKeyValues() {
	seq := it.SliceValues([]tuple.Tuple2[int, string]{
		tuple.New2(0, "a"),
		tuple.New2(1, "b"),
		tuple.New2(2, "c"),
	})
	for k, v := range it.KeyValues(seq) {
		fmt.Printf("%d: %s\n", k, v)
	}

	// Output:
	// 0: a
	// 1: b
	// 2: c
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

func ExampleUnzip() {
	seq1, seq2 := it.Unzip(it.SliceIter([]string{"a", "b", "c"}))

	for v := range seq1 {
		fmt.Println(v)
	}
	for v := range seq2 {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
	// a
	// b
	// c
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

func ExampleAny() {
	fmt.Println(it.Any(it.Range(5), 4))
	fmt.Println(it.Any(it.Range(5), 5))
	// Output:
	// true
	// false
}

func ExampleAny2() {
	m := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4"}
	fmt.Println(it.Any2(it.MapIter(m), 1, "1"))
	fmt.Println(it.Any2(it.MapIter(m), 1, "2"))
	// Output:
	// true
	// false
}

func ExampleAnyBy() {
	fmt.Println(it.AnyBy(it.Range(5), func(v int) bool { return v > 3 }))
	fmt.Println(it.AnyBy(it.Range(5), func(v int) bool { return v > 5 }))
	// Output:
	// true
	// false
}

func ExampleAnyBy2() {
	m := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4"}
	fmt.Println(it.AnyBy2(it.MapIter(m), func(_ int, v string) bool { return v == "1" }))
	fmt.Println(it.AnyBy2(it.MapIter(m), func(_ int, v string) bool { return v == "5" }))
	// Output:
	// true
	// false
}

func ExampleAll() {
	fmt.Println(it.All(it.Range(1), 0))
	fmt.Println(it.All(it.Range(5), 0))
	// Output:
	// true
	// false
}

func ExampleAll2() {
	m := map[int]string{0: "0"}
	fmt.Println(it.All2(it.MapIter(m), 0, "0"))
	fmt.Println(it.All2(it.MapIter(m), 0, "1"))
	// Output:
	// true
	// false
}

func ExampleAllBy() {
	fmt.Println(it.AllBy(it.Range(5), func(v int) bool { return v < 5 }))
	fmt.Println(it.AllBy(it.Range(5), func(v int) bool { return v < 4 }))
	// Output:
	// true
	// false
}

func ExampleAllBy2() {
	m := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4"}
	fmt.Println(it.AllBy2(it.MapIter(m), func(_ int, v string) bool { return v < "5" }))
	fmt.Println(it.AllBy2(it.MapIter(m), func(_ int, v string) bool { return v < "4" }))
	// Output:
	// true
	// false
}

func ExampleMax() {
	max := it.Max(it.SliceValues([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}))
	fmt.Println(max.GetOrZero())
	// Output:
	// 9
}

func ExampleMin() {
	min := it.Min(it.SliceValues([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}))
	fmt.Println(min.GetOrZero())
	// Output:
	// 1
}
