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

## Support of range over func

utility functions for iterators are available with [the it package](https://pkg.go.dev/github.com/shogo82148/hi/it).

```go
import (
    "github.com/shogo82148/hi/it"
)
```

## Reference

### SliceValues, MapKeys, MapValues, ChanValues

[SliceValues](https://pkg.go.dev/github.com/shogo82148/hi/it#SliceValues), [MapKeys](https://pkg.go.dev/github.com/shogo82148/hi/it#MapKeys), [MapValues](https://pkg.go.dev/github.com/shogo82148/hi/it#MapValues) and [ChanValues](https://pkg.go.dev/github.com/shogo82148/hi/it#ChanValues) create
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

[SliceIter](https://pkg.go.dev/github.com/shogo82148/hi/it#SliceIter) and [MapIter](https://pkg.go.dev/github.com/shogo82148/hi/it#MapIter) create iterators
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

[Range](https://pkg.go.dev/github.com/shogo82148/hi/it#Range) creates an iterator that returns a sequence of integers.

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

[Enumerate](https://pkg.go.dev/github.com/shogo82148/hi/it#Enumerate) returns an iterator that returns (index, value) tuples from seq.

```go
seq := it.Enumerate(it.SliceValues([]string{"a", "b", "c"}))
// it.SliceIter([]string{"a", "b", "c"})
```

### Cycle, Cycle2

[Cycle](https://pkg.go.dev/github.com/shogo82148/hi/it#Cycle) makes an iterator returning elements from the iterable and saving a copy of each.

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

[Tee](https://pkg.go.dev/github.com/shogo82148/hi/it#Tee) makes n iterators from seq.

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

[it.Zip](https://pkg.go.dev/github.com/shogo82148/hi/it#Zip) converts a pair of [iter.Seq](https://pkg.go.dev/iter#Seq) to [iter.Seq2](https://pkg.go.dev/iter#Seq2).

```go
it.Zip(it.Range(5), it.SliceValues([]string{"zero", "one", "two"}))
// it.SliceIter([]string{"zero", "one", "two"})
```

### ZipLongest

[ZipLongest](https://pkg.go.dev/github.com/shogo82148/hi/it#ZipLongest) is the same as Zip, but the returned iterator continues until the longest iterator is exhausted.
If a shorter iterator is exhausted, it is filled with zero values.

```go
it.Zip(it.Range(5), it.SliceValues([]string{"zero", "one", "two"}))
// it.SliceIter([]string{"zero", "one", "two", "", ""})
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

Zip2,...,Zip16 on iterators:

```go
it.Zip2(it.Range(3), it.SliceValues([]string{"zero", "one", "two"}))
// it.SliceValues([]tuple.Tuple2{
//     tuple.New2(0, "zero"),
//     tuple.New2(1, "one"),
//     tuple.New2(2, "two"),
// })
```

### Unzip

[it.Unzip](https://pkg.go.dev/github.com/shogo82148/hi/it#Unzip) converts a [iter.Seq2](https://pkg.go.dev/iter#Seq2) to a pair of [iter.Seq](https://pkg.go.dev/iter#Seq).

```go
it.Unzip(it.SliceIter([]string{"zero", "one", "two"}))
// it.Range(3), it.SliceValues([]string{"zero", "one", "two"})
```

### Unzip2,... , Unzip16

```go
hi.Unzip2([]tuple.Tuple2{tuple.New2(1, "one"), tuple.New2(2, "two"), tuple.New2(3, "three")})
// []int{1, 2, 3}, []string{"one", "two", "three"}
```

Unzip2,..., Unzip16 on iterators:

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

[Map](https://pkg.go.dev/github.com/shogo82148/hi/it#Map) on iterators:

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

[Filter](https://pkg.go.dev/github.com/shogo82148/hi/it#Filter) on iterators:

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

[Chunk](https://pkg.go.dev/github.com/shogo82148/hi/it#Chunk) on iterators:

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

[Slice](https://pkg.go.dev/github.com/shogo82148/hi/it#Slice) on iterators:

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

[Chain](https://pkg.go.dev/github.com/shogo82148/hi/it#Chain) on iterators:

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

[Compress](https://pkg.go.dev/github.com/shogo82148/hi/it#Compress) on iterators:

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

[DropWhile](https://pkg.go.dev/github.com/shogo82148/hi/it#DropWhile) on iterators:

```go
it.DropWhile(it.SliceValues([]int{1, 4, 6, 4, 1}), func(v int) bool { return v < 5 })
// it.SliceValues([]int{6, 4, 1})

it.DropWhile2(it.SliceIter([]int{1, 4, 6, 4, 1}), func(_, v int) bool { return v < 5 })
// it.Zip(it.SliceValues([]int{2, 3, 4}), it.SliceValues([]int{6, 4, 1}))
```

### Pairwise

[Pairwise](https://pkg.go.dev/github.com/shogo82148/hi#Pairwise) makes a slice of all adjacent pairs of elements.

```go
l := hi.Pairwise([]int{1, 2, 3, 4, 5})
// []tuple.Tuple2{tuple.New2(1, 2), tuple.New2(2, 3), tuple.New2(3, 4), tuple.New2(4, 5)}
```

[Pairwise](https://pkg.go.dev/github.com/shogo82148/hi/it#Pairwise) on iterators:

```go
seq := it.Pairwise(it.Range(5))
// it.SliceValues([]tuple.Tuple2{tuple.New2(1, 2), tuple.New2(2, 3), tuple.New2(3, 4), tuple.New2(4, 5)})

seq := it.Pairwise2(it.SliceIter([]string{"zero", "one", "two"}))
// it.Zip(
//     it.SliceValues([]tuple.Tuple2{tuple.New2(0, 1), tuple.New2(1, 2)})
//     it.SliceValues([]tuple.Tuple2{tuple.New2("zero", "one"), tuple.New2("onw", "two")})
// )
```

### TakeWhile

[TakeWhile](https://pkg.go.dev/github.com/shogo82148/hi#TakeWhile) returns a slice of the longest prefix of elements that satisfy predicate f.

```go
l := hi.TakeWhile([]int{1, 4, 6, 4, 1}, func(_, val int) bool {
    return v < 5
})
// []int{1, 4}
```

[TakeWhile](https://pkg.go.dev/github.com/shogo82148/hi/it#TakeWhile) on iterators:

```go
it.TakeWhile(it.SliceValues([]int{1, 4, 6, 4, 1}), func(v int) bool { return v < 5 })
// it.SliceValues([]int{1, 4})

it.TakeWhile2(it.SliceIter([]int{1, 4, 6, 4, 1}), func(_, v int) bool { return v < 5 })
// it.Zip(it.SliceValues([]int{0, 1}), it.SliceValues([]int{1, 4}))
```

### Repeat

[it.Repeat](https://pkg.go.dev/github.com/shogo82148/hi/it#Repeat) returns an infinitely continuing iterator.

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

[RepeatN](https://pkg.go.dev/github.com/shogo82148/hi/it#RepeatN) on iterators:

```go
it.RepeatN("hello", 3)
// it.SliceValues([]string{"hello", "hello", "hello"})
```

### Count

```go
hi.Count([]int{1, 2, 3, 4, 5}, 3)
// 1
```

[Count](https://pkg.go.dev/github.com/shogo82148/hi/it#Count) on iterators:

```go
it.Count(it.Range(5), 3)
// 1
```

### CountBy

```go
numberOfEven := hi.CountBy([]int{1, 2, 3, 4, 5}, func(_, v int) bool { return v % 2 == 0 })
// 2
```

[CountBy](https://pkg.go.dev/github.com/shogo82148/hi/it#CountBy) on iterators:

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

[Any](https://pkg.go.dev/github.com/shogo82148/hi/it#Any) on iterators

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

[AnyBy](https://pkg.go.dev/github.com/shogo82148/hi/it#AnyBy) on iterators:

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

[All](https://pkg.go.dev/github.com/shogo82148/hi/it#All) on iterators

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

[AllBy](https://pkg.go.dev/github.com/shogo82148/hi/it#AllBy) on iterators:

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

[Max](https://pkg.go.dev/github.com/shogo82148/hi/it#Max) on iterators:

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

[Min](https://pkg.go.dev/github.com/shogo82148/hi/it#Min) on iterators:

```go
it.Min(it.SliceValues([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}))
// optional.New(1)
```

### Shuffle

[Shuffle](https://pkg.go.dev/github.com/shogo82148/hi#Shuffle) returns a shuffled copy of a.

```go
s := hi.Shuffle([]int{1, 2, 3, 4, 5})
// []int{2, 3, 1, 5, 4}
```

[Shuffle](https://pkg.go.dev/github.com/shogo82148/hi/it#Shuffle) on iterators:

```go
s := it.Shuffle(it.Range(5))
// it.SliceValues([]int{1, 2, 0, 4, 3})

s := it.Shuffle(it.SliceIter([]string{"zero", "one", "two"}))
// it.SliceValues([]tuple.Tuple2{tuple.New(1, "one"), tuple.New(0, "zero"), tuple.New(2, "two")})
```

### Sample

```go
v := hi.Sample([]int{1, 2, 3, 4, 5})
// optional.New(2)
```

[Sample](https://pkg.go.dev/github.com/shogo82148/hi/it#Sample) on iterators:

```go
v := it.Sample(it.Range(5))
// optional.New(5)

v := it.Sample2(it.SliceIter([]string{"zero", "one", "two", "three", "four"}))
// optional.New(tuple.New2(2, "two"))
```

### SampleN

```go
s := hi.Sample([]int{1, 2, 3, 4, 5}, 3)
// []int{4, 5, 2}
```

[SampleN](https://pkg.go.dev/github.com/shogo82148/hi/it#SampleN) on iterators:

```go
s := it.SampleN(it.Range(5))
// []int{3, 4, 1}

s := it.SampleN2(it.SliceIter([]string{"zero", "one", "two", "three", "four"}), 3)
// []tuple.Tuple2{tuple.New2(3, "three"), tuple.New2(4, "four"), tuple.New2(1, "one")}
```

### Ptr

```go
p := hi.Ptr("hello")
// *p = "hello"
```
