package hi_test

import (
	"fmt"

	"github.com/shogo82148/hi"
)

func ExampleNegate() {
	f := func(v int) bool { return v > 3 }
	g := hi.Negate(f)

	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %t, %t\n", i, f(i), g(i))
	}

	// Output:
	// 0: false, true
	// 1: false, true
	// 2: false, true
	// 3: false, true
	// 4: true, false
}

func ExampleNegate2() {
	f := func(a, b int) bool { return a > b }
	g := hi.Negate2(f)

	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %t, %t\n", i, f(i, 3), g(i, 3))
	}

	// Output:
	// 0: false, true
	// 1: false, true
	// 2: false, true
	// 3: false, true
	// 4: true, false
}

// Example of Negate with Filter.
func ExampleNegate2_filter() {
	filter := func(_, v int) bool { return v > 3 }
	negate := hi.Negate2(filter)
	fmt.Println(hi.Filter([]int{1, 2, 3, 4, 5}, negate))

	// Output:
	// [1 2 3]
}
