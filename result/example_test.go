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
