package list_test

import (
	"fmt"

	"github.com/shogo82148/hi/list"
	"github.com/shogo82148/hi/tuple"
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

func ExampleUnzip2() {
	var l list.List[tuple.Tuple2[int, string]]
	l.PushBack(tuple.New2(1, "one"))
	l.PushBack(tuple.New2(2, "two"))
	l.PushBack(tuple.New2(3, "three"))

	l1, l2 := list.Unzip2(&l)
	for e := l1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// Output:
	// 1
	// 2
	// 3
	// one
	// two
	// three
}

func ExampleFilter() {
	var l list.List[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	m := list.Filter(&l, func(_, v int) bool { return v > 3 })
	for e := m.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// Output:
	// 4
	// 5
}

func ExampleCount() {
	var l list.List[int]
	l.PushBack(3)
	l.PushBack(1)
	l.PushBack(4)
	l.PushBack(1)
	l.PushBack(5)
	l.PushBack(9)
	l.PushBack(2)
	l.PushBack(6)
	l.PushBack(5)
	l.PushBack(3)
	l.PushBack(5)

	cnt := list.Count(&l, 5)
	fmt.Println(cnt)
	// Output:
	// 3
}

func ExampleCountBy() {
	var l list.List[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	cnt := list.CountBy(&l, func(_, v int) bool { return v > 3 })
	fmt.Println(cnt)
	// Output:
	// 2
}

func ExampleAny() {
	var l list.List[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	fmt.Println(list.Any(&l, 5))
	fmt.Println(list.Any(&l, 6))
	// Output:
	// true
	// false
}

func ExampleAnyBy() {
	var l list.List[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	fmt.Println(list.AnyBy(&l, func(_, v int) bool { return v > 3 }))
	fmt.Println(list.AnyBy(&l, func(_, v int) bool { return v > 5 }))
	// Output:
	// true
	// false
}

func ExampleAll() {
	var l list.List[int]
	l.PushBack(5)
	l.PushBack(5)
	l.PushBack(5)
	l.PushBack(5)
	l.PushBack(5)

	fmt.Println(list.All(&l, 5))
	fmt.Println(list.Any(&l, 6))
	// Output:
	// true
	// false
}

func ExampleAllBy() {
	var l list.List[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	fmt.Println(list.AllBy(&l, func(_, v int) bool { return v > 0 }))
	fmt.Println(list.AllBy(&l, func(_, v int) bool { return v > 3 }))
	// Output:
	// true
	// false
}
