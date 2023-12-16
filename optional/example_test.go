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

func ExampleFilter() {
	filter := func(v int) bool { return v > 20 }

	v := optional.New(10).Filter(filter)
	fmt.Println(v.Valid())

	w := optional.New(30).Filter(filter)
	fmt.Println(w.Valid())

	// Output:
	// false
	// true
}

func ExampleFilterFalse() {
	filter := func(v int) bool { return v > 20 }

	v := optional.New(10).FilterFalse(filter)
	fmt.Println(v.Valid())

	w := optional.New(30).FilterFalse(filter)
	fmt.Println(w.Valid())

	// Output:
	// true
	// false
}

func ExampleMap() {
	v := optional.New(10)
	w := optional.Map(v, func(v int) string { return fmt.Sprintf("%x", v) })
	fmt.Println(w.GetOrZero())
	// Output:
	// a
}
