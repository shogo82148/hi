package hi_test

import (
	"fmt"

	"github.com/shogo82148/hi"
	"github.com/shogo82148/hi/tuple"
)

func ExampleZip2() {
	ints := []int{1, 2, 3, 4, 5}
	strings := []string{"one", "two", "three", "four", "five"}
	slice := hi.Zip2(ints, strings)
	for _, v := range slice {
		fmt.Println(v)
	}
	// Output:
	// (1, one)
	// (2, two)
	// (3, three)
	// (4, four)
	// (5, five)
}

func ExampleUnzip2() {
	slice := []tuple.Tuple2[int, string]{
		tuple.New2(1, "one"),
		tuple.New2(2, "two"),
		tuple.New2(3, "three"),
		tuple.New2(4, "four"),
		tuple.New2(5, "five"),
	}
	s1, s2 := hi.Unzip2(slice)
	for _, v := range s1 {
		fmt.Println(v)
	}
	for _, v := range s2 {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// one
	// two
	// three
	// four
	// five
}

func ExampleMap() {
	input := []int{1, 2, 3, 4, 5}
	output := hi.Map(input, func(_, v int) string {
		return fmt.Sprintf("(%d)", v)
	})
	for _, v := range output {
		fmt.Println(v)
	}
	// Output:
	// (1)
	// (2)
	// (3)
	// (4)
	// (5)
}

func ExampleFilter() {
	input := []int{1, 2, 3, 4, 5}
	output := hi.Filter(input, func(_, v int) bool {
		return v%2 == 0
	})
	for _, v := range output {
		fmt.Println(v)
	}
	// Output:
	// 2
	// 4
}

func ExampleFilterFalse() {
	input := []int{1, 2, 3, 4, 5}
	output := hi.FilterFalse(input, func(_, v int) bool {
		return v%2 == 0
	})
	for _, v := range output {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 3
	// 5
}

func ExampleGroupBy() {
	input := []int{1, 2, 3, 4, 5}
	output := hi.GroupBy(input, func(_, v int) int {
		return v % 2
	})
	for k, v := range output {
		fmt.Println(k, v)
	}
	// Unordered output:
	// 0 [2 4]
	// 1 [1 3 5]
}

func ExampleChunk() {
	input := []int{1, 2, 3, 4, 5}
	output := hi.Chunk(input, 2)
	for _, v := range output {
		fmt.Println(v)
	}
	// Output:
	// [1 2]
	// [3 4]
	// [5]
}

func ExampleSlice() {
	input := []int{1, 2, 3, 4, 5}
	output := hi.Slice(input, 1, 3, 1)
	fmt.Println(output)
	// Output:
	// [2 3]
}

func ExampleSlice_NegativeIndex() {
	input := []int{1, 2, 3, 4, 5}

	// slices last 2 elements
	output := hi.Slice(input, -2, len(input), 1)
	fmt.Println(output)

	// Output:
	// [4 5]
}

func ExampleSlice_NegativeStep() {
	input := []int{1, 2, 3, 4, 5}
	output := hi.Slice(input, 3, 1, -1)
	fmt.Println(output)
	// Output:
	// [4 3]
}

func ExampleChain() {
	input1 := []int{1, 2, 3}
	input2 := []int{4, 5, 6}
	output := hi.Chain(input1, input2)
	for _, v := range output {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExampleRepeatN() {
	output := hi.RepeatN("hello", 3)
	for _, v := range output {
		fmt.Println(v)
	}
	// Output:
	// hello
	// hello
	// hello
}

func ExampleCount() {
	input := []int{1, 2, 3, 4, 5}
	fmt.Println(hi.Count(input, 3))
	// Output:
	// 1
}

func ExampleCountBy() {
	input := []int{1, 2, 3, 4, 5}
	fmt.Println(hi.CountBy(input, func(_, v int) bool { return v%2 == 0 }))
	// Output:
	// 2
}

func ExampleAny() {
	input := []int{1, 2, 3, 4, 5}
	fmt.Println(hi.Any(input, 5))
	fmt.Println(hi.Any(input, 6))
	// Output:
	// true
	// false
}

func ExampleAnyBy() {
	input := []int{1, 2, 3, 4, 5}
	fmt.Println(hi.AnyBy(input, func(_, v int) bool { return v > 3 }))
	fmt.Println(hi.AnyBy(input, func(_, v int) bool { return v > 10 }))
	// Output:
	// true
	// false
}

func ExampleAll() {
	fmt.Println(hi.All([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(hi.All([]int{5, 5, 5, 5, 5}, 5))
	// Output:
	// false
	// true
}

func ExampleAllBy() {
	fmt.Println(hi.AllBy([]int{1, 2, 3, 4, 5}, func(_, v int) bool { return v > 3 }))
	fmt.Println(hi.AllBy([]int{4, 5, 6, 7, 8}, func(_, v int) bool { return v > 3 }))
	// Output:
	// false
	// true
}

func ExampleMax() {
	max := hi.Max([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5})
	fmt.Println(max.Get())
	// Output:
	// 9
}

func ExampleMin() {
	min := hi.Min([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5})
	fmt.Println(min.Get())
	// Output:
	// 1
}
