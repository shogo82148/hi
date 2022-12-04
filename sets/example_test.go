package sets_test

import (
	"fmt"

	"github.com/shogo82148/hi/sets"
)

func ExampleSet_Count() {
	s := sets.New(1, 2, 3, 4, 5)
	cnt := s.Count(func(v int) bool { return v > 3 })
	fmt.Println(cnt)
	// Output:
	// 2
}

func ExampleSet_Any() {
	s := sets.New(1, 2, 3, 4, 5)
	fmt.Println(s.All(func(v int) bool { return v > 0 }))
	fmt.Println(s.Any(func(v int) bool { return v > 3 }))
	fmt.Println(s.Any(func(v int) bool { return v > 5 }))
	// Output:
	// true
	// true
	// false
}

func ExampleSet_All() {
	s := sets.New(1, 2, 3, 4, 5)
	fmt.Println(s.All(func(v int) bool { return v > 0 }))
	fmt.Println(s.All(func(v int) bool { return v > 3 }))
	fmt.Println(s.Any(func(v int) bool { return v > 5 }))
	// Output:
	// true
	// false
	// false
}
