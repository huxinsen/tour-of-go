package main

import (
	"fmt"
	"math"
	"runtime"
)

// Go has only one looping construct, the for loop.

// The init statement will often be a short variable declaration, and the
// variables declared there are visible only in the scope of the for statement.

// Note: Unlike other languages like C, Java, or JavaScript there are no
// parentheses surrounding the three components of the for statement and
// the braces { } are always required.

// The init and post statements are optional. And for is Go's "while".
// If you omit the loop condition it loops forever,
// so an infinite loop is compactly expressed.

func main() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i) // 0 1 2 3 4 5 6 7 8 9
			continue LABEL
		}
	}

	fmt.Println(sqrt(2), sqrt(-4)) // 1.4142135623730951 2i

	fmt.Println(
		pow(3, 2, 10), // 9
		pow(3, 3, 20), // 20
	)

	printOS() // Go runs on windows.
}

func sqrt(x float64) string {
	// Go's if statements are like its for loops; the expression need not be
	// surrounded by parentheses ( ) but the braces { } are required.
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// Sprint formats using the default formats for its operands and
	// returns the resulting string.
	return fmt.Sprint(math.Sqrt(x))
}

// Like for, the if statement can start with a short
// statement to execute before the condition.

// Variables declared by the statement are only in scope until the end of the if.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func printOS() {
	fmt.Print("Go runs on ")
	// Go only runs the selected case, not all the cases that follow.
	// In effect, the break statement that is needed at the end of each case
	// is provided automatically in Go.

	// Go's switch cases need not be constants,
	// and the values involved need not be integers.

	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
