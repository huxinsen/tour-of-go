package main

import (
	"fmt"
	"sync"
	"time"
)

// A `goroutine` is a lightweight thread managed by the Go runtime.

// go f(x, y, z)
// starts a new goroutine running f(x, y, z)
// The evaluation of f, x, y, and z happens in the current goroutine
// and the execution of f happens in the new goroutine.

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sayMain() {
	go say("world")
	say("hello")
}

// Channels are a typed conduit through which you can send and receive values
// with the channel operator, <-.
// ch <- v    // Send v to channel ch.
// v := <-ch  // Receive from ch, and assign value to v.
// (The data flows in the direction of the arrow.)

// Like maps & slices, channels must be created before use: ch := make(chan int)
// By default, sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks
// or condition variables.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sumMain() {
	// The example code sums the numbers in a slice, distributing the work between
	// two goroutines. Once both goroutines have completed their computation,
	// it calculates the final result.

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

// Channels can be buffered. Provide the buffer length as the second argument
// to make to initialize a buffered channel:
// ch := make(chan int, 100)

// Sends to a buffered channel block only when the buffer is full.
// Receives block when the buffer is empty.

// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3 // fatal error: all goroutines are asleep - deadlock!
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// A sender can `close` a channel to indicate that no more values will be sent.
// Receivers can test whether a channel has been closed by assigning a second
// parameter to the receive expression:
// v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.

// Note: Only the sender should close a channel, never the receiver.
// Sending on a closed channel will cause a panic.
func fibonacci(n int, c chan int) {
	defer close(c)
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
}

func fibonacciMain() {
	// The loop for i := range c receives values from
	// the channel repeatedly until it is closed.
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}
}

// Another note: Channels aren't like files; you don't usually need to close
// them. Closing is only necessary when the receiver must be told there are
// no more values coming, such as to terminate a range loop.

// The select statement lets a goroutine wait on multiple communication
// operations. A select blocks until one of its cases can run,
// then it executes that case. It chooses one at random if multiple are ready.
func selectMain() {
	// Tick is a convenience wrapper for NewTicker
	// providing access to the ticking channel only.
	tick := time.Tick(100 * time.Millisecond)
	// After waits for the duration to elapse and then
	// sends the current time on the returned channel.
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// sync.Mutex

// We've seen how channels are great for communication among goroutines.
// But what if we don't need communication? What if we just want to make sure
// only one goroutine can access a variable at a time to avoid conflicts?

// This concept is called mutual exclusion, and the conventional name for the
// data structure that provides it is mutex.

// Go's standard library provides mutual exclusion with sync.Mutex and
// its two methods: Lock, Unlock

// We can define a block of code to be executed in mutual exclusion by
// surrounding it with a call to Lock and Unlock as shown on the Inc method.

// We can also use defer to ensure the mutex
// will be unlocked as in the Value method.

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func mutexMain() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

func main() {
	sayMain()
	sumMain()
	fibonacciMain()
	selectMain()
	mutexMain()
}
