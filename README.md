# go-generic

[![Go Reference](https://pkg.go.dev/badge/github.com/NublyBR/go-generic.svg)](https://pkg.go.dev/github.com/NublyBR/go-generic)
[![Go Report Card](https://goreportcard.com/badge/github.com/NublyBR/go-generic)](https://goreportcard.com/report/github.com/NublyBR/go-generic)

This repository contains a collection of Go utility functions and types built with Go's generics.

## Example usage

```go
package main

import (
	"fmt"

	"github.com/NublyBR/go-generic"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Filter

	evens := generic.Filter(func(i int) bool {
		return i%2 == 0
	}, ints...)

	fmt.Println(evens) // [2 4 6 8 10 12 14 16 18 20]

	// Map

	mult := generic.Map(func(i int) int {
		return i * 2
	}, ints...)

	fmt.Println(mult) // [2 4 6 8 10 12 14 16 18 20 22 24 26 28 30 32 34 36 38 40]
}
```