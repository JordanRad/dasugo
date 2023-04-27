# DasuGo

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/dasugo)](https://goreportcard.com/report/github.com/yourusername/dasugo)

DasuGo is a Go utility library inspired by JavaScript's functional programming methods, such as map, reduce, filter, some, and intersect. It is written in Go and makes use of Go generics.

## Features

- A collection of functional programming methods inspired by JavaScript
- Written in Go with support for Go generics
- Easy to use and extend
- Well-tested and production-ready

## Installation

To use DasuGo in your Go project, simply run:

```bash
go get github.com/JordanRad/dasugo
```

## Usage

Here are some examples of how you can use DasuGo in your Go code:

```go
import (
    "fmt"
    "github.com/JordanRad/dasugo"
)

func main() {
    // Map example
    numbers := []int{1, 2, 3, 4, 5}
    doubledNumbers := dasugo.Map(numbers, func(n int) int {
        return n * 2
    })
    fmt.Println(doubledNumbers) // Output: [2 4 6 8 10]

    // Filter example
    words := []string{"foo", "bar", "baz", "qux", "quux"}
    filteredWords := dasugo.Filter(words, func(s string) bool {
        return len(s) < 4
    })
    fmt.Println(filteredWords) // Output: [foo bar baz]

    // Reduce example
    numbers := []int{1, 2, 3, 4, 5}
    sum := dasugo.Reduce(numbers, func(acc, n int) int {
        return acc + n
    }, 0)
    fmt.Println(sum) // Output: 15

    // Some example
    numbers := []int{1, 2, 3, 4, 5}
    hasEvenNumbers := dasugo.Some(numbers, func(n int) bool {
        return n%2 == 0
    })
    fmt.Println(hasEvenNumbers) // Output: true

    // Contains example
    numbers := []int{1, 2, 3, 4, 5}
    containsThree := dasugo.Contains(numbers, 3)
    fmt.Println(containsThreee) // Output: true

    // Intersect example
    numbers1 := []int{1, 2, 3, 4, 5}
    numbers2 := []int{4, 5, 6, 7, 8}
    commonNumbers := dasugo.Intersect(numbers1, numbers2)
    fmt.Println(commonNumbers) // Output: [4 5]
}
```

## Contributing

We welcome contributions to DasuGo! To contribute, simply fork this repository, make your changes, and submit a pull request. Before submitting your pull request, please make sure that your code passes the tests and that you have added tests for any new functionality.

## License

DasuGo is licensed under the MIT License. See the LICENSE file for more information.