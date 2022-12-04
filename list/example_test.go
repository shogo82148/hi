package list_test

import (
	"fmt"

	"github.com/shogo82148/hi/list"
)

func ExampleZip2() {
	var l1 list.List[int]
	var l2 list.List[string]
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l2.PushBack("one")
	l2.PushBack("two")
	l2.PushBack("three")

	l := list.Zip2(&l1, &l2)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// Output:
	// (1, one)
	// (2, two)
	// (3, three)
}
