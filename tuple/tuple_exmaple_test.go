package tuple_test

import (
	"fmt"

	"github.com/shogo82148/hi/tuple"
)

func ExampleNew2() {
	t := tuple.New2(1, "one")
	fmt.Println(t)
	// Output:
	// (1, one)
}
