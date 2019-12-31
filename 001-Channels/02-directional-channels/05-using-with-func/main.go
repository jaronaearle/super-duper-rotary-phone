package main

import "fmt"

func main() {
	// channels are reference type, references underlying data
	c := make(chan int)
	// send
	go foo(c)

	//receive
	bar(c)

	fmt.Println("about to exit")

}

// send
func foo(c chan<- int) {
	// puts 42 onto c
	c <- 42
}

//receive
func bar(c <-chan int) {
	// pulls a value off of c
	fmt.Println(<-c)
}
