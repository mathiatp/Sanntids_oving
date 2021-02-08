// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing() {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000000; j++ {
		i++
	}
}

func decrementing() {
	//TODO: decrement i 1000000 times
	for k := 0; k < 1000000; k++ {
		i--
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// Try doing the exercise both with and without it!

	// TODO: Spawn both functions as goroutines
	go incrementing()
	go decrementing()

	// We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
	// We'll come back to using channels in Exercise 2. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}

//What is happening? The two functions is running parallell. That means there is a chance that incrementing will read i=1, 
//then decrementing read i = 1, since incrementing has not written i=2 back. Then incrementing will continoue i++ and then write i = 2 back,
// but at the same time decrementing is computing with i = 1 and i-- to i = 0 and write back i = 0; We have now lost the incrementation in addition to running decrementing one time.
// This can happen several times and that is why we get seamingly random numbers of positive or negative sign.
