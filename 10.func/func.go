package main

import (
	"fmt"
	"math"
)

func main() {
	functionValues()

	f := fibonacci()
	for i := 0; i < 6; i++ {
		fmt.Println(f()) // 0, 1, 1, 2, 3, 5
	}

	paramsAndEffects()
}

func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13

	fmt.Println(compute(hypot))    // 5
	fmt.Println(compute(math.Pow)) // 81
}

// Functions are values too. They can be passed around just like other values.
// Function values may be used as function arguments and return values.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Go functions may be closures. A closure is a function value that references
// variables from outside its body. The function may access and assign to the
// referenced variables; in this sense the function is "bound" to the variables.

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	pre, next := 0, 1
	return func() int {
		result := pre
		pre, next = next, pre+next
		return result
	}
}

func paramsAndEffects() {
	a, b, c := 1, 2, 3
	funcA(a, b, c)       // [4 5 6]
	fmt.Println(a, b, c) // 1 2 3

	s := []int{1, 2, 3}
	funcB(s)       // [4 5 6]
	fmt.Println(s) // [4 5 6]

	n := 1
	funcC(&n)      // 2
	fmt.Println(n) // 2
}

// Variadic function parameters
func funcA(s ...int) {
	s[0] = 4
	s[1] = 5
	s[2] = 6
	fmt.Println(s)
}

func funcB(s []int) {
	s[0] = 4
	s[1] = 5
	s[2] = 6
	fmt.Println(s)
}

func funcC(n *int) {
	*n++
	fmt.Println(*n)
}
