# hi - utility functions for collections

[![Go Reference](https://pkg.go.dev/badge/github.com/shogo82148/hi.svg)](https://pkg.go.dev/github.com/shogo82148/hi)

hi is a Lodash-style Go library based on Go 1.18+ Generics inspired by [lo](https://github.com/samber/lo).

## Usage

Install the package:

```
go get github.com/shogo82148/hi
```

and import it:

```go
import (
    "github.com/shogo82148/hi"
)
```

## Experimental support of range over func

Experimental support of [range over func](https://github.com/golang/go/issues/61405) is available from Go 1.22.

```
GOEXPERIMENT=rangefunc go1.22.0 run main.go
```

utility functions for iterators are available with `it` package.

```go
import (
    "github.com/shogo82148/hi/it"
)
```

## Reference

### SliceValues, MapKeys, MapValues, ChanValues

Experimental: `SliceValues`, `MapKeys`, `MapValues` and `ChanValues` creates
iterators from slices, maps and channels.

```go
for v := range it.SliceValues([]string{"a", "b", "c"}) {
    fmt.Println(v)
}

// Output:
// a
// b
// c
```

### SliceIter, MapIter

Experimental: `SliceIter` and `MapIter` create iterators
that return key-value pairs.

```go
for i, v := range it.SliceIter([]string{"a", "b", "c"}) {
    fmt.Printf("%d: %s\n", i, v)
}

// Output:
// 0: a
// 1: b
// 2: c
```

### Range

Experimental: `Range` creates an iterator that returns a sequence of integers.

```go
for v := range it.Range(5) {
    fmt.Println(v)
}

// Output:
// 0
// 1
// 2
// 3
// 4
```

### Enumerate

Experimental: Enumerate returns an iterator that returns (index, value) tuples from seq.

```go
seq := it.Enumerate(it.SliceValues([]string{"a", "b", "c"}))
// it.SliceIter([]string{"a", "b", "c"})
```

### Cycle, Cycle2

Experimental: Cycle makes an iterator returning elements from the iterable and saving a copy of each.

```go
seq := it.Cycle(it.Range(3))
// it.SliceValues([]int{0, 1, 2, 0, 1, 2, 0, ...})

seq := it.Cycle2(it.SliceIter([]string{"zero", "one", "two"}))
// it.Zip(
//     it.SliceValues([]int{0, 1, 2, 0, 1, 2, 0, ...}),
//     it.SliceValues([]string{"zero", "one", "two", "zero", "one", "two", "zero", ...}),
// )
```

### Tee, Tee2

Experimental: Tee makes n iterators from seq.

```go
s := it.Tee(it.Range(3), 2)
// []iter.Seq[int]{it.Range(3), it.Range(3)}
```

```go
s := it.Tee2(it.SliceIter([]string{"zero", "one", "two"}), 2)
// []iter.Seq2[int, string]{
//     it.SliceIter([]string{"zero", "one", "two"}),
//     it.SliceIter([]string{"zero", "one", "two"}),
// }
```

### Zip

Experimental: `it.Zip` converts a pair of `iter.Seq` to `iter.Seq2`.

```go
it.Zip(it.Range(3), it.SliceValues([]string{"zero", "one", "two"}))
// it.SliceIter([]string{"zero", "one", "two"})
```

### Zip2,... ,Zip16

```go
hi.Zip2([]int{1, 2, 3}, []string{"one", "two", "three"})
// []tuple.Tuple2{
//     tuple.New2(1, "one"),
//     tuple.New2(2, "two"),
//     tuple.New2(3, "three"),
// }
```

Experimental: zip on iterators

```go
it.Zip2(it.Range(3), it.SliceValues([]string{"zero", "one", "two"}))
// it.SliceValues([]tuple.Tuple2{
//     tuple.New2(0, "zero"),
//     tuple.New2(1, "one"),
//     tuple.New2(2, "two"),
// })
```

### Unzip

Experimental: `it.Unzip` converts a `iter.Seq2` to a pair of `iter.Seq`.

```go
it.Unzip(it.SliceIter([]string{"zero", "one", "two"}))
// it.Range(3), it.SliceValues([]string{"zero", "one", "two"})
```

### Unzip2,... , Unzip16

```go
hi.Unzip2([]tuple.Tuple2{tuple.New2(1, "one"), tuple.New2(2, "two"), tuple.New2(3, "three")})
// []int{1, 2, 3}, []string{"one", "two", "three"}
```

Experimental: Unzip2,..., Unzip16 on iterators

```go
seq := it.Unzip(it.SliceValues([]tuple.Tuple2{
    tuple.New2(0, "zero"),
    tuple.New2(1, "one"),
    tuple.New2(2, "two"),
}))
// it.Range(3), it.SliceValues([]string{"zero", "one", "two"})
```

### Map

```go
hi.Map([]int{1, 2, 3, 4, 5}, func(_, v int) string {
    return fmt.Sprintf("(%d)", v)
})
// []string{"(1)", "(2)", "(3)", "(4)", "(5)"}
```

Experimental: `Map` on iterators

```go
it.Map(it.Range(5), func(v int) string {
    return fmt.Sprintf("(%d)", v)
})
// it.SliceValues([]string{"(0)", "(1)", "(2)", "(3)", "(4)"})
```

### Filter

```go
even := hi.Filter([]int{1, 2, 3, 4, 5}, func(_, v int) bool {
    return v%2 == 0
})
// []int{2, 4}
```

Experimental: `Filter` on iterators

```go
even := it.Filter(it.Range(5), func(v int) bool {
    return v%2 == 0
})
// it.SliceValues([]int{0, 2, 4})
```

### Chunk

```go
hi.Chunk([]int{1, 2, 3, 4, 5}, 2)
// [][]int{{1, 2}, {3, 4}, {5}}
```

Experimental: `Chunk` on iterators

```go
it.Chunk(it.Range(5), 2)
// it.SliceValues([][]int{{0, 1}, {2, 3}, {4}})
```

### Slice

Python-like slice.

```go
hi.Slice([]int{0, 1, 2, 3, 4}, 1, 3, 1)
// []int{1, 2}
```

Experimental: `Slice` on iterators

```go
it.Slice(it.Range(5), 1, 3, 1)
// it.SliceValues([]int{1, 2})

it.Slice2(it.SliceIter([]string{"a", "b", "c", "d", "e", "f", "g"}), 2, 4, 1)
// it.Zip(it.SliceValues([]int{2, 3}), it.SliceValues([]string{"c", "d"}))
```

### Chain

[Chain](https://pkg.go.dev/github.com/shogo82148/hi#Chain) returns a slice of all elements from all slices.

```go
hi.Chain([]int{1, 2, 3}, []int{4, 5, 6})
// []int{1, 2, 3, 4, 5, 6}
```

Experimental: `Chain` on iterators

```go
it.Chain(it.Range(3), it.Range(3))
// it.SliceValues([]int{0, 1, 2, 0, 1, 2})

it.Chain2(
    it.SliceIter([]string{"zero", "one", "two"}),
    it.SliceIter([]string{"null", "eins", "zwei"}),
)
// it.Zip([]int{0, 1, 2, 0, 1, 2}, []string{"zero", "one", "two", "null", "eins", "zwei"})
```

### Compress

[Compress](https://pkg.go.dev/github.com/shogo82148/hi#Compress) makes a slice of the elements in a for which the corresponding element in selectors is true.

```go
hi.Compress([]int{1, 2, 3, 4, 5}, []bool{true, false, true, false, true})
// []int{1, 3, 5}
```

Experimental: `Compress` on iterators

```go
it.Compress(it.Range(5), it.SliceValues([]bool{true, false, true, false, true}))
// it.SliceValues([]int{0, 2, 4})

it.Compress2(it.SliceIter([]string{"zero", "one", "two"}), it.SliceValues([]bool{true, false, true}))
// it.Zip(it.SliceValues([]int{0, 2}), it.SliceValues([]string{"zero", "two"}))
```

### DropWhile

[DropWhile](https://pkg.go.dev/github.com/shogo82148/hi#DropWhile) returns a slice of the remaining elements after dropping the longest prefix of elements that satisfy predicate.

```go
l := hi.DropWhile([]int{1, 4, 6, 4, 1}, func(_, val int) bool {
    return v < 5
})
// []int{6, 4, 1}
```

Experimental: `DropWhile` on iterators

```go
it.DropWhile(it.SliceValues([]int{1, 4, 6, 4, 1}), func(v int) bool { return v < 5 })
// it.SliceValues([]int{6, 4, 1})

it.DropWhile2(it.SliceIter([]int{1, 4, 6, 4, 1}), func(_, v int) bool { return v < 5 })
// it.Zip(it.SliceValues([]int{2, 3, 4}), it.SliceValues([]int{6, 4, 1}))
```

### TakeWhile

[TakeWhile](https://pkg.go.dev/github.com/shogo82148/hi#TakeWhile) returns a slice of the longest prefix of elements that satisfy predicate f.

```go
l := hi.TakeWhile([]int{1, 4, 6, 4, 1}, func(_, val int) bool {
    return v < 5
})
// []int{1, 4}
```

Experimental: `TakeWhile` on iterators

```go
it.TakeWhile(it.SliceValues([]int{1, 4, 6, 4, 1}), func(v int) bool { return v < 5 })
// it.SliceValues([]int{1, 4})

it.TakeWhile2(it.SliceIter([]int{1, 4, 6, 4, 1}), func(_, v int) bool { return v < 5 })
// it.Zip(it.SliceValues([]int{0, 1}), it.SliceValues([]int{1, 4}))
```

### Repeat

Experimental: `Range` on iterators

```go
for v := range it.Repeat("hello") {
    fmt.Println(v)
}
// hello
// hello
// hello
// ... (infinite)
```

### RepeatN

```go
hi.RepeatN("hello", 3)
// []string{"hello", "hello", "hello"}
```

Experimental: `RepeatN` on iterators

```go
it.RepeatN("hello", 3)
// it.SliceValues([]string{"hello", "hello", "hello"})
```

### Count

```go
hi.Count([]int{1, 2, 3, 4, 5}, 3)
// 1
```

Experimental: `Count` on iterators

```go
it.Count(it.Range(5), 3)
// 1
```

### CountBy

```go
numberOfEven := hi.CountBy([]int{1, 2, 3, 4, 5}, func(_, v int) bool { return v % 2 == 0 })
// 2
```

Experimental: `CountBy` on iterators

```go
numberOfEven := it.CountBy(it.Range(5), func(_, v int) bool { return v % 2 == 0 })
// 3
```

### Any

[Any](https://pkg.go.dev/github.com/shogo82148/hi#Any) returns whether l has value at least one.

```go
hi.Any([]int{1, 2, 3, 4, 5}, 5)
// true
```

Experimental: `Any` on iterators

```go
it.Any(it.Range(5), 4)
// true

it.Any2(it.SliceIter([]int{1, 2, 3, 4, 5}), 4, 5)
// true
```

### AnyBy

[AnyBy](https://pkg.go.dev/github.com/shogo82148/hi#AnyBy) returns whether a has an element for that f returns true.

```go
hi.AnyBy([]int{1, 2, 3, 4, 5}, func(_, v int) bool { return v > 3 })
// true
```

Experimental: `AnyBy` on iterators

```go
it.AnyBy(it.Range(5), func(v int) bool { return v > 3 })
// true

it.AnyBy2(it.SliceIter([]int{1, 2, 3, 4, 5}), func(_, v int) bool { return v > 3 })
// true
```

### All

[All](https://pkg.go.dev/github.com/shogo82148/hi#All) returns whether all elements of a are value.

```go
hi.All([]int{1, 2, 3, 4, 5}, 5)
// false

hi.All([]int{5, 5, 5, 5, 5}, 5)
// true
```

Experimental: `All` on iterators

```go
it.All(it.Range(1), 0)
// true

it.Any2(it.SliceIter([]int{1}), 0, 1)
// true
```

### AllBy

[AllBy](https://pkg.go.dev/github.com/shogo82148/hi#AllBy) returns whether f returns true for all elements in a.

```go
hi.AllBy([]int{1, 2, 3, 4, 5}, func(_, v int) bool { return v > 3 })
// false

hi.AllBy([]int{4, 5, 6, 7, 8}, func(_, v int) bool { return v > 3 })
// true
```

Experimental: `AllBy` on iterators

```go
it.AllBy(it.Range(5), func(v int) bool { return v < 5 })
// true

it.AllBy2(it.MapIter(map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4"}), func(_ int, v string) bool { return v < "5" })
// true
```

### Max

[Max](https://pkg.go.dev/github.com/shogo82148/hi#Max) returns the maximum element of arguments.

```go
hi.Max([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5})
// optional.New(9)
```

Experimental: `Max` on iterators

```go
it.Max(it.SliceValues([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}))
// optional.New(9)
```

### Min

[Min](https://pkg.go.dev/github.com/shogo82148/hi#Min) returns the minimum element of arguments.

```go
hi.Min([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5})
// optional.New(1)
```

Experimental: `Min` on iterators

```go
it.Min(it.SliceValues([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}))
// optional.New(1)
```

### Ptr

```go
p := hi.Ptr("hello")
// *p = "hello"
```
