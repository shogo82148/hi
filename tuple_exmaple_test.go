package hi

import "fmt"

func ExampleNewTuple2() {
	t := NewTuple2(1, "one")
	fmt.Println(t)
	// Output:
	// (1, one)
}
