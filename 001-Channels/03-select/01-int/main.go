package main

import "fmt"

func main() {
	// channels for evens, odds, and quit
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	// calls send as a goroutine to run on its own
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) { // takes in 3 input only channels
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i // puts evens on e
		} else {
			o <- i // puts odds on o
		}
	}
	close(e)
	close(o)
	q <- 0 // puts 0, quit value on q
}

func receive(e, o, q <-chan int) { // takes in 3 output channels only
	for {
		select { // select statement listens for and pulls values off of available channels and can do something with the value pulled
		case v := <-e:
			fmt.Println("from the eve channel: ", v)
		case v := <-o:
			fmt.Println("from the odd channel: ", v)
		case v := <-q:
			fmt.Println("from the quit channel: ", v)
			return
		}
	}
}
