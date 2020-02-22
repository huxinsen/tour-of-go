package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Connector interface {
	Connect()
}

// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.

type PhoneConnector struct {
	name string
}

// This method means type `PhoneConnector` implements the interface `Connector`,
// but we don't need to explicitly declare that it does so.

func (pc PhoneConnector) Connect() {
	fmt.Println("Connected: ", pc.name)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	pc := PhoneConnector{"PhoneConnector"}
	pc.Connect()   // Connected:  PhoneConnector
	Disconnect(pc) // Disconnected:  PhoneConnector

	// Under the hood, interface values can be thought of as a tuple of a value
	// and a concrete type: (value, type)
	// An interface value holds a value of a specific underlying concrete type.
	// Calling a method on an interface value executes the method of the same
	// name on its underlying type.
	var c Connector
	c = Connector(pc)
	c.Connect() // Connected:  PhoneConnector

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f                // a MyFloat implements Abser
	fmt.Println(a.Abs()) // 1.4142135623730951
	a = &v               // a *Vertex implements Abser
	fmt.Println(a.Abs()) // 5

	// A nil interface value holds neither value nor concrete type.
	var i I
	describe(i) //(<nil>, <nil>)

	// Calling a method on a nil interface is a run-time error because there is no
	// type inside the interface tuple to indicate which concrete method to call.
	// i.M() // runtime error

	// Note that an interface value that holds a nil concrete value
	// is itself non-nil.
	var t *T
	i = t
	describe(i) // (<nil>, *main.T)

	// If the concrete value inside the interface itself is nil,
	// the method will be called with a nil receiver.
	i.M() // <nil>

	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T)
	i.M()       // hello

	// The interface type that specifies zero methods is known as
	// the empty interface: interface{}
	// An empty interface may hold values of any type.
	var ei interface{}
	describeAll(ei) // (<nil>, <nil>)

	ei = 42
	describeAll(ei) // (42, int)

	ei = "hello"
	describeAll(ei) // (hello, string)
}

// A type assertion provides access to an interface value's underlying concrete
// value. t := i.(T)
// This statement asserts that the interface value i holds the concrete type T
// and assigns the underlying T value to the variable t.

// If i does not hold a T, the statement will trigger a panic.

// To test whether an interface value holds a specific type, a type assertion
// can return two values: the underlying value and a boolean value that reports
// whether the assertion succeeded. t, ok := i.(T)
// If i holds a T, then t will be the underlying value and ok will be true.
// If not, ok will be false and t will be the zero value of type T,
// and no panic occurs.

// Empty interfaces are used by code that handles values of unknown type.

func Disconnect(connector interface{}) {
	// A type switch is like a regular switch statement, but the cases in a type
	// switch specify types (not values), and those values are compared against
	// the type of the value held by the given interface value.
	switch v := connector.(type) {
	case PhoneConnector:
		// here v has type PhoneConnector
		fmt.Println("Disconnected: ", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeAll(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
