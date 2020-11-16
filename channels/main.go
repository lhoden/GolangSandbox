package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	wg.Done()
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int, 10) // A channel of the int type, buffering for 10 items
	for i := 0; i < 10; i++ {
		wg.Add(1) // need one of these for every go routine
		go foo(fooVal, i)
	}
	wg.Wait() // tells the channels to stay open until all the go routines are finished
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}
	// At this point, we're running all of the go routines, but
	// if the program closes too quickly and the routines are
	// out running but don't have a chance to come back
	// Go lets us use syncing to synchronize all the Go routines
	// which gives us the fastest possible result and performance
	// We do this by establishing wait groups
	// Channels are something that Go is very well known for! Maximizing performance!
}
