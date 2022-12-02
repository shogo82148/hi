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
