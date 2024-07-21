package sets_test

import (
	"fmt"

	"github.com/shogo82148/hi/sets"
)

func ExampleFilter() {
	s := sets.New(1, 2, 3, 4, 5)
	u := sets.Filter(s, func(v int) bool { return v > 3 })

	for v := range u {
		fmt.Println(v)
	}
	// Unordered output:
	// 4
	// 5
}

func ExampleFilterFalse() {
	s := sets.New(1, 2, 3, 4, 5)
	u := sets.FilterFalse(s, func(v int) bool { return v > 3 })

	for v := range u {
		fmt.Println(v)
	}
	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleCountBy() {
	s := sets.New(1, 2, 3, 4, 5)
	cnt := sets.CountBy(s, func(v int) bool { return v > 3 })
	fmt.Println(cnt)
	// Output:
	// 2
}

func ExampleAnyBy() {
	s := sets.New(1, 2, 3, 4, 5)
	fmt.Println(sets.AnyBy(s, func(v int) bool { return v > 0 }))
	fmt.Println(sets.AnyBy(s, func(v int) bool { return v > 3 }))
	fmt.Println(sets.AnyBy(s, func(v int) bool { return v > 5 }))
	// Output:
	// true
	// true
	// false
}

func ExampleAllBy() {
	s := sets.New(1, 2, 3, 4, 5)
	fmt.Println(sets.AllBy(s, func(v int) bool { return v > 0 }))
	fmt.Println(sets.AllBy(s, func(v int) bool { return v > 3 }))
	fmt.Println(sets.AllBy(s, func(v int) bool { return v > 5 }))
	// Output:
	// true
	// false
	// false
}
