package hi_test

import (
	"fmt"

	"github.com/shogo82148/hi"
)

func ExampleNewTuple2() {
	t := hi.NewTuple2(1, "one")
	fmt.Println(t)
	// Output:
	// (1, one)
}
