// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)

func number_server(add <-chan int, sub chan<- int, read <-chan int) {
	var number = 0

	for {
		select {
		case num := <-add:
			number += num
		case sub <- number:
		case <-read:
			return
		}
	}
}

func incrementer(add chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add <- 1
	}
	// Signal that the goroutine is finished
	finished <- true

}

func decrementer(sub chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		sub <- -1
	}
	// Signal that the goroutine is finished
	finished <- true

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Construct the required channels
	CH_add := make(chan int)
	CH_sub := make(chan int)
	CH_read := make(chan int)

	CH_incr_finished := make(chan bool)
	CH_decr_finished := make(chan bool)


	// Spawn the required goroutines
	go number_server(CH_add, CH_sub, CH_read)
	go incrementer(CH_add,CH_incr_finished)
	go decrementer(CH_add,CH_decr_finished)

	// Block on finished from both "worker" goroutines
	<- CH_incr_finished
	<- CH_decr_finished

	fmt.Println("The magic number is:", <-CH_sub)
	CH_read <- 0
}
