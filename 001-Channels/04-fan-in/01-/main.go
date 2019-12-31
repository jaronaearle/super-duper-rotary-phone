package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	// sends
	go send(even, odd)

	// receives
	// pulls values from even/odd and puts into fanIn
	go receive(even, odd, fanIn)

	// shows values in fanIn
	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// populates even and odd with numbers
func send(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

// pulls numbers from even/odd and populates fanIn
func receive(e, o <-chan int, f chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() { //pulls evens and puts into f
		for v := range e {
			f <- v
		}
		wg.Done()
	}()

	go func() { // pulls odds and puts into f
		for v := range o {
			f <- v
		}
		wg.Done()
	}()

	wg.Wait() // waits and closes
	close(f)
}
