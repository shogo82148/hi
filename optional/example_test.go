package optional_test

import (
	"fmt"

	"github.com/shogo82148/hi/optional"
	"github.com/shogo82148/hi/tuple"
)

func ExampleZip2() {
	opt := optional.Zip2(optional.New(1), optional.New("one"))
	fmt.Println(opt.GetOrZero())
	// Output:
	// (1, one)
}

func ExampleUnzip2() {
	o1, o2 := optional.Unzip2(optional.New(tuple.New2(1, "one")))
	fmt.Println(o1.GetOrZero())
	fmt.Println(o2.GetOrZero())
	// Output:
	// 1
	// one
}
