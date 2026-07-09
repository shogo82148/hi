package hi_test

import (
	"fmt"
)

func ExamplePtr() {
	p := new("hello")
	fmt.Println(*p)

	// Output:
	// hello
}
