package main

import "fmt"

func main() {
	c := make(chan<- int, 2) // send only channel

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("--------")
	fmt.Printf("%T\n", c)
}
