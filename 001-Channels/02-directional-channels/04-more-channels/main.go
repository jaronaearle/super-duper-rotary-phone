package main

import (
	"fmt"
)

func main() {
	c := make(chan int)    // send/receive channel
	cr := make(<-chan int) // receive channel
	cs := make(chan<- int) // send channel

	fmt.Println("-----------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

}
