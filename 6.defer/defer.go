package main

import (
	"fmt"
)

func main() {
	a() // Func a
	b() // Panic in b Recover in b
	c() // Func c
	d()
	// Func d
	// closure i =  4
	// closure i =  4
	// closure i =  4
	// closure i =  4
	// closure_fix i =  0
	// closure_fix i =  1
	// closure_fix i =  2
	// closure_fix i =  3
	// defer_closure i =  4
	// defer i =  3
	// defer_closure i =  4
	// defer i =  2
	// defer_closure i =  4
	// defer i =  1
	// defer_closure i =  4
	// defer i =  0
	fmt.Println(e()) // 2
}

func a() {
	fmt.Println("Func a")
}

// A defer statement defers the execution of a function
// until the surrounding function returns.

// Defer is commonly used to simplify functions
// that perform various clean-up actions.

func b() {
	defer func() {
		// Recover is a built-in function that regains control of a panicking
		// goroutine. Recover is only useful inside deferred functions.
		// During normal execution, a call to recover will return nil and have
		// no other effect. If the current goroutine is panicking, a call to recover
		// will capture the value given to panic and resume normal execution.

		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Recover in b")
		}
	}()

	// Panic is a built-in function that stops the ordinary flow of control and
	// begins panicking. When the function F calls panic, execution of F stops,
	// any deferred functions in F are executed normally, and then F returns to
	// its caller. To the caller, F then behaves like a call to panic. The process
	// continues up the stack until all functions in the current goroutine have
	// returned, at which point the program crashes. Panics can be initiated by
	// invoking panic directly. They can also be caused by runtime errors, such
	//  as out-of-bounds array accesses.
	panic("Panic in b")
}

func c() {
	fmt.Println("Func c")
}

func d() {
	fmt.Println("Func d")
	var fs = [4]func(){}
	var fs2 = [4]func(){}

	for i := 0; i < 4; i++ {
		// A deferred function's arguments are evaluated
		// when the defer statement is evaluated.

		// Deferred function calls are executed in Last In First Out order
		// after the surrounding function returns.
		defer fmt.Println("defer i = ", i)

		defer func() {
			fmt.Println("defer_closure i = ", i)
		}()

		fs[i] = func() { fmt.Println("closure i = ", i) }

		fs2[i] = func(i int) func() {
			return func() {
				fmt.Println("closure_fix i = ", i)
			}
		}(i)
	}

	for _, f := range fs {
		f()
	}

	for _, f := range fs2 {
		f()
	}
}

// Deferred functions may read and assign to the
// returning function's named return values.
func e() (i int) {
	defer func() { i++ }()
	return 1
}
