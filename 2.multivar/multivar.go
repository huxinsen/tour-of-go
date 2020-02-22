package main

import (
	"fmt"
	"math/cmplx"
)

// The var statement declares a list of variables.
// A var statement can be at package or function level.

// If an initializer is present, the type can be omitted.
var (
	name             = "daniel"
	hobby            = "coding"
	z     complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var a, b = 1, 2

	// Inside a function, the := short assignment statement
	// can be used in place of a var declaration with implicit type.
	c := 3

	// Hello, I am daniel. I like coding.
	fmt.Println("Hello, I am " + name + ". I like " + hobby + ".")
	fmt.Println(a)                           // 1
	fmt.Println(b)                           // 2
	fmt.Println(c)                           // 3
	fmt.Printf("Type: %T Value: %v\n", z, z) // Type: complex128 Value: (2+3i)
}

// Go's basic types are
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32
//      // represents a Unicode code point
// float32 float64
// complex64 complex128

// The int, uint, and uintptr types are usually 32 bits wide on
// 32-bit systems and 64 bits wide on 64-bit systems. When you
// need an integer value you should use int unless you have a
// specific reason to use a sized or unsigned integer type.

// Zero values
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.
