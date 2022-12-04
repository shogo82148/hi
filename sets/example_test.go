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
