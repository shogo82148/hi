package result_test

import (
	"fmt"

	"github.com/shogo82148/hi/result"
	"github.com/shogo82148/hi/tuple"
)

func ExampleUnzip2() {
	ret := result.OK(tuple.New2(1, "one"))
	r1, r2 := result.Unzip2(ret)
	fmt.Println(r1.Get())
	fmt.Println(r2.Get())
	// Output:
	// 1
	// one
}

func ExampleFilter() {
	filter := func(v int) bool { return v > 20 }

	v := result.OK(10).Filter(filter)
	fmt.Println(v.Err())

	w := result.OK(30).Filter(filter)
	fmt.Println(w.Err())

	// Output:
	// result: filtered
	// <nil>
}

func ExampleFilterFalse() {
	filter := func(v int) bool { return v > 20 }

	v := result.OK(10).FilterFalse(filter)
	fmt.Println(v.Err())

	w := result.OK(30).FilterFalse(filter)
	fmt.Println(w.Err())

	// Output:
	// <nil>
	// result: filtered
}
