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

Experimental: `Range` returns an iterator that generates.

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
