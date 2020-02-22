package main

// The first statement in a Go source file must be package name.
// Executable commands must always use package main.

// By convention, the package name is the same as
// the last element of the import path.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("hello, world\n") // hello, world

	// In Go, a name is exported if it begins with a capital letter.
	fmt.Println(math.Pi) // 3.141592653589793

	a, b := swap("hello", "world")
	fmt.Println(a, b) // world hello

	fmt.Println(split(17)) // 7 10
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named.
// If so, they are treated as variables defined at the top of the function.
// These names should be used to document the meaning of the return values.

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
