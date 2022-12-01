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
