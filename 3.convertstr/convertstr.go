package main

import (
	"fmt"
	"strconv"
)

// Unlike in C, in Go assignment between items of
// different type requires an explicit conversion.

func main() {
	var a int = 65

	// The expression T(v) converts the value v to the type T.
	b := string(a)

	// Itoa is equivalent to FormatInt(int64(i), 10).
	c := strconv.Itoa(a)

	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	d, _ := strconv.Atoi(c)

	fmt.Printf("Type: %T Value: %v\n", b, b) // Type: string Value: A
	fmt.Printf("Type: %T Value: %v\n", c, c) // Type: string Value: 65
	fmt.Printf("Type: %T Value: %v\n", d, d) // Type: int Value: 65
}
