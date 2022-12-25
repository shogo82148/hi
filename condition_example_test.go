package hi_test

import (
	"fmt"

	"github.com/shogo82148/hi"
)

func ExampleIf_even() {
	v := 124
	fmt.Println(hi.If(v%2 == 0, "even", "odd"))
	// Output:
	// even
}

func ExampleIf_odd() {
	v := 125
	fmt.Println(hi.If(v%2 == 0, "even", "odd"))
	// Output:
	// odd
}

func ExampleSwitch() {
	v := 2
	fmt.Println(hi.Switch[string](v).
		Case(1, "one").
		Case(2, "two").
		Case(3, "three").
		Default("four or greater"))
	// Output:
	// two
}
