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

func ExampleMax() {
	max := hi.Max(3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5)
	fmt.Println(max.GetOrZero())
	// Output:
	// 9
}

func ExampleMaxBy() {
	max := hi.MaxBy(
		func(s string) int { return len(s) },
		"Can", "I", "find", "a", "trick", "recalling", "Pi", "easily",
	)
	fmt.Println(max.GetOrZero())
	// Output:
	// recalling
}

func ExampleMin() {
	min := hi.Min(3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5)
	fmt.Println(min.GetOrZero())
	// Output:
	// 1
}

func ExampleMinBy() {
	min := hi.MinBy(
		func(s string) int { return len(s) },
		"Can", "I", "find", "a", "trick", "recalling", "Pi", "easily",
	)
	fmt.Println(min.GetOrZero())
	// Output:
	// I
}

func ExampleMinMax() {
	min, max := hi.MinMax(3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5)
	fmt.Println("min:", min.GetOrZero())
	fmt.Println("max:", max.GetOrZero())
	// Output:
	// min: 1
	// max: 9
}

func ExampleMinMaxBy() {
	min, max := hi.MinMaxBy(
		func(s string) int { return len(s) },
		"Can", "I", "find", "a", "trick", "recalling", "Pi", "easily",
	)
	fmt.Println("min:", min.GetOrZero())
	fmt.Println("max:", max.GetOrZero())
	// Output:
	// min: I
	// max: recalling
}

func ExampleSum() {
	sum := hi.Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(sum.Get())
	// Output:
	// 55
}

func ExampleReduce() {
	sum := hi.Reduce([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(_ int, sum int, item int) int {
		return sum + item
	}, 0)
	fmt.Println(sum)
	// Output:
	// 55
}
