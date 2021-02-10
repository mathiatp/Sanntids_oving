package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		ch <- i
	}

}

func consumer(ch <-chan int) {

	time.Sleep(1 * time.Second)
	for {
		i := <-ch
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	// TODO: make a bounded buffer
	ch_transport := make(chan int, 5)

	go consumer(ch_transport)
	go producer(ch_transport)

	select {}
}
