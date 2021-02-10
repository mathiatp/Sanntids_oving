// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)

func number_server(ch_add <-chan int, ch_sub <-chan int, ch_read chan<- int) {
	number := 0
	for {
		select {
		case <-ch_add:
			number++
		case <-ch_sub:
			number--
		case ch_read <- number:
			break
		}
	}
}

func increment(ch_add chan<- int, ch_finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		ch_add <- 1
	}
	// Signal that the goroutine is finished
	ch_finished <- true
}

func decrement(ch_sub chan<- int, ch_finished chan<- bool) {
	for j := 0; j < 1000001; j++ {
		ch_sub <- 1
	}
	// Signal that the goroutine is finished
	ch_finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Construct the required channels
	ch_add := make(chan int)
	ch_sub := make(chan int)
	ch_read := make(chan int)
	ch_finished := make(chan bool, 2)

	// Spawn the required goroutines
	go number_server(ch_add, ch_sub, ch_read)
	go decrement(ch_sub, ch_finished)
	go increment(ch_add, ch_finished)

	<-ch_finished
	<-ch_finished

	// Block on finished from both "worker" goroutines
	fmt.Println("The magic number is:", <-ch_read)

}
