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
