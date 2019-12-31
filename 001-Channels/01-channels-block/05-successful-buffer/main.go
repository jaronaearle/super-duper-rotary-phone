package main

import "fmt"

func main() {
	c := make(chan int, 5)

	c <- 42
	c <- 43
	c <- 7

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
