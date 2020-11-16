package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup: ", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}

func defering() {
	defer fmt.Println("Done!")
	fmt.Println("Doing some stuff, who knows what?")

	//an example with a loop
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("friend!")
	//wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()

	defering()
}
