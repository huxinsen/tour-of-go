package main

import (
	"fmt"
)

// Constants are declared like variables, but with the const keyword.
// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax.

// Storage Unit
const (
	// The iota identifier resets to 0 whenever the word const appears
	// in the source code and increments after each const specification.

	// "iota" is the letter of the Greek alphabet.
	// It is typical for the math notations:
	// - as iterator in sums and algorithms
	// - as subscript index
	// - for imaginary part of complex numbers

	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	const World = "Storage Unit"
	fmt.Println("Hello", World) // Hello Storage Unit

	fmt.Println(B)  // 1
	fmt.Println(KB) // 1024
	fmt.Println(MB) // 1.048576e+06
	fmt.Println(GB) // 1.073741824e+09
	fmt.Println(TB) // 1.099511627776e+12
	fmt.Println(PB) // 1.125899906842624e+15
	fmt.Println(EB) // 1.152921504606847e+18
	fmt.Println(ZB) // 1.1805916207174113e+21
	fmt.Println(YB) // 1.2089258196146292e+24

	// An untyped constant takes the type needed by its context.
	const Big = 1 << 100
	fmt.Println("needFloat", needFloat(Big)) // needFloat 1.2676506002282295e+29
}

func needFloat(x float64) float64 {
	return x * 0.1
}
