package main

import (
	"fmt"
)

// A struct is a collection of fields.

type Human struct {
	Name    string
	Gender  string
	Age     int
	Contact struct {
		Address, Tel string
	}
}

type Teacher struct {
	Human
	Subject string
}

type Student struct {
	Human
	Class string
}

type SuperType struct {
	Name string
}

type SubType struct {
	Super SuperType
	Name  string
}

func main() {
	// Go has pointers. A pointer holds the memory address of a value.
	// The type *T is a pointer to a T value. Its zero value is nil.
	i := 42

	// The & operator generates a pointer to its operand.
	p := &i // point to i

	// The * operator denotes the pointer's underlying value.
	// read i through the pointer
	fmt.Println(*p) // 42

	*p = 21        // set i through the pointer
	fmt.Println(i) // 21

	var t Teacher
	// Struct fields are accessed using a dot.
	t.Name = "Tom"
	t.Gender = "male"
	t.Age = 30
	t.Contact.Address = "Beijing"
	t.Contact.Tel = "18833445566"
	t.Subject = "CS"

	var s Student
	s.Name = "Daniel"
	s.Gender = "male"
	s.Age = 18
	s.Contact.Address = "Beijing"
	s.Contact.Tel = "19911223344"
	s.Class = "One"

	fmt.Println(t) // {{Tom male 30 {Beijing 18833445566}} CS}
	fmt.Println(s) // {{Daniel male 18 {Beijing 19911223344}} One}

	// You can list just a subset of fields by using the Name: syntax.
	// (And the order of named fields is irrelevant.)
	a := SubType{Name: "name-sub", Super: SuperType{Name: "name-super"}}
	fmt.Println(a.Name, a.Super.Name) // name-sub name-super
	fmt.Println(a)                    // {{name-super} name-sub}

	// Struct fields can be accessed through a struct pointer.
	// To access the field X of a struct when we have the struct pointer p we
	// could write (*p).X. However, that notation is cumbersome, so the language
	// permits us instead to write just p.X, without the explicit dereference.
	pa := &a
	fmt.Println(pa.Name) // name-sub
}
