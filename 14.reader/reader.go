package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// The io package specifies the io.Reader interface,
// which represents the read end of a stream of data.

// The io.Reader interface has a Read method:

// func (T) Read(b []byte) (n int, err error)
// Read populates the given byte slice with data and
// returns the number of bytes populated and an error
// value. It returns an io.EOF error when the stream ends.

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	switch {
	case x >= 'A' && x <= 'M':
		fallthrough
	case x >= 'a' && x <= 'm':
		x = x + 13
	case x >= 'N' && x <= 'Z':
		fallthrough
	case x >= 'n' && x <= 'z':
		x = x - 13
	}
	return x
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	// NewReader returns a new Reader reading from s. It is similar
	// to bytes.NewBufferString but more efficient and read-only.
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	// n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
	// b[:n] = "Hello, R"
	// n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
	// b[:n] = "eader!"
	// n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
	// b[:n] = ""

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r13 := rot13Reader{s}

	// func Copy(dst Writer, src Reader) (written int64, err error)
	// Copy copies from src to dst until either EOF is reached
	// on src or an error occurs. It returns the number of bytes
	// copied and the first error encountered while copying, if any.
	io.Copy(os.Stdout, &r13) // You cracked the code!
}
