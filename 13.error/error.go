package main

import (
	"fmt"
)

// One of the most ubiquitous interfaces is Stringer defined by the fmt package.

// type Stringer interface {
// 	String() string
// }
// A Stringer is a type that can describe itself as a string. The fmt package
// (and many others) look for this interface to print values.

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}

// The error type is a built-in interface similar to fmt.Stringer:

// type error interface {
// 	Error() string
// }
// (As with fmt.Stringer, the fmt package looks for
// the error interface when printing values.)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	// Newton's method: https://en.wikipedia.org/wiki/Newton%27s_method
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip) // loopback: 127.0.0.1 googleDNS: 8.8.8.8
	}

	fmt.Println(Sqrt(2))  // 1.414213562373095 <nil>
	fmt.Println(Sqrt(-2)) // 0 cannot Sqrt negative number: -2
}
