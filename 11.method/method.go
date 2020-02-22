package main

import (
	"fmt"
	"math"
)

// Go does not have classes. However, you can define methods on types.

// A method is a function with a special receiver argument. The receiver appears
// in its own argument list between the func keyword and the method name.

type Vertex struct {
	X, Y float64
}

// In this example, the Abs method has a receiver of type Vertex named v.
// Remember: a method is just a function with a receiver argument.

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can only declare a method with a receiver whose type is
// defined in the same package as the method. You cannot declare
// a method with a receiver whose type is defined in another
// package (which includes the built-in types such as int).

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Methods with pointer receivers can modify the value to which the receiver
// points (as Scale does here). Since methods often need to modify their
// receiver, pointer receivers are more common than value receivers.

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// Before scaling: {X:3 Y:4}, Abs: 5
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) // 1.4142135623730951

	// Functions with a pointer(value) argument must take a pointer(value),
	// while methods with pointer(value) receivers take either a value
	// or a pointer as the receiver when they are called.

	v.Scale(10)
	// After scaling: {X:30 Y:40}, Abs: 50
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

	// There are two reasons to use a pointer receiver. The first is so that the
	// method can modify the value that its receiver points to. The second is to
	// avoid copying the value on each method call. This can be more efficient if
	// the receiver is a large struct, for example.
}
