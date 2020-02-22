package main

import (
	"fmt"
)

func main() {
	// The type [n]T is an array of n values of type T.
	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // Hello World
	fmt.Println(a)          // [Hello World]

	// An array has a fixed size. A slice, on the other hand, is a
	// dynamically-sized, flexible view into the elements of an array.
	// The type []T is a slice with elements of type T.
	b := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// A slice is formed by specifying two indices, a low and high bound,
	// separated by a colon: a[low : high]
	// This selects a half-open range which includes the first element,
	// but excludes the last one. The following expression creates a slice which
	// includes elements 5 through 7 of b
	s1 := b[5:8]

	// The default is zero for the low bound and the length of the slice for the
	// high bound. These slice expressions are equivalent:
	// b[0:10] b[:10] b[0:] b[:]
	s2 := b[:]

	printSlice("s1", s1) // s1 len=3 cap=5 [5 6 7]

	// A slice does not store any data, it just describes a section of an
	// underlying array. Changing the elements of a slice modifies the
	// corresponding elements of its underlying array.
	s1[0] += 10
	printSlice("s1", s1) // s1 len=3 cap=5 [15 6 7]
	fmt.Println("b", b)  // b [0 1 2 3 4 15 6 7 8 9]

	// Other slices that share the same underlying array will see those changes.
	printSlice("s2:", s2) // s2: len=10 cap=10 [0 1 2 3 4 15 6 7 8 9]

	// You can extend a slice's length by re-slicing it,
	// provided it has sufficient capacity.
	s2 = s2[:3]
	printSlice("s2:", s2) // s2: len=3 cap=10 [0 1 2]

	// The make function allocates a zeroed array and
	// returns a slice that refers to that array
	s3 := make([]int, 3, 6)
	printSlice("s3", s3) // s3 len=3 cap=6 [0 0 0]

	// func append(s []T, vs ...T) []T
	// The first parameter s of append is a slice of type T,
	// and the rest are T values to append to the slice.

	// The resulting value of append is a slice containing all
	// the elements of the original slice plus the provided values.

	// If the backing array of s is too small to fit all the given
	// values a bigger array will be allocated. The returned slice
	// will point to the newly allocated array.
	s3 = append(s3, 1, 2, 3)
	fmt.Printf("s3: %v &s3: %p\n", s3, s3) // s3: [0 0 0 1 2 3] &s3: 0xc00000a330

	s3 = append(s3, 1, 2, 3)
	// s3: [0 0 0 1 2 3 1 2 3] &s3: 0xc0000480c0
	fmt.Printf("s3: %v &s3: %p\n", s3, s3)

	// A slice literal is like an array literal without the length.
	s4 := []int{1, 2, 3}
	copy(s4[1:3], b[8:10])
	printSlice("s4:", s4) // s4: len=3 cap=3 [1 8 9]

	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var s5 []int
	printSlice("s5:", s5) // s5: len=0 cap=0 []
	if s5 == nil {
		fmt.Println("nil!") // nil!
	}

	// Slices can contain any type, including other slices.
}

func printSlice(s string, x []int) {
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array,
	// counting from the first element in the slice.
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
