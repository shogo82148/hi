package hi_test

import (
	"fmt"

	"github.com/shogo82148/hi"
)

func ExamplePtr() {
	p := hi.Ptr("hello")
	fmt.Println(*p)

	// Output:
	// hello
}
