package main

import (
	"fmt"
	"strings"
)

var pow = []int{1, 2, 4, 8}

func main() {
	// The range form of the for loop iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration. The
	// first is the index, and the second is a copy of the element at that index.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v) // 2**0 = 1 2**1 = 2 2**2 = 4 2**3 = 8
	}

	// You can skip the index or value by assigning to _.
	// for i, _ := range pow
	// for _, value := range pow

	// A map maps keys to values.
	// The zero value of a map is nil. A nil map has no keys,
	// nor can keys be added.
	var m1 map[string]int
	m1 = wordCount("I Love You! I Love You! I Love You!")
	fmt.Println(m1) // map[I:3 Love:3 You!:3]

	// Delete an element:
	delete(m1, "Love")
	fmt.Println(m1) // map[I:3 You!:3]

	// Insert an element:
	m1["Love"] = 10000
	// Retrieve an element and update an element:
	m1["I"] = m1["Love"]
	m1["You!"] = m1["Love"]
	fmt.Println(m1) // map[I:10000 Love:10000 You!:10000]

	// Test that a key is present with a two-value assignment:
	// elem, ok = m[key]
	// If key is in m, ok is true. If not, ok is false. If key is not in the map,
	// then elem is the zero value for the map's element type.
	v, ok := m1["Love"]
	fmt.Println("The value:", v, "Present?", ok) // The value: 10000 Present? true

	// Map literals are like struct literals, but the keys are required.
	m2 := map[int]string{
		5: "我",
		2: "爱",
		0: "你",
	}
	fmt.Println(m2) // map[0:你 2:爱 5:我]

}

func wordCount(s string) map[string]int {
	// The make function returns a map of the given type,
	// initialized and ready for use.
	result := make(map[string]int)

	// Fields splits the string s around each instance of one or more consecutive
	// white space characters, as defined by unicode. IsSpace, returning a slice
	// of substrings of s or an empty slice if s contains only white space.
	arr := strings.Fields(s)
	for _, v := range arr {
		result[v]++
	}
	return result
}
