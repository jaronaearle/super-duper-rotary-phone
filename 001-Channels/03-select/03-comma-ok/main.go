package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 7
		c <- 9
		close(c)
	}()

	// pulls value 7 from c, returns true ok
	v, ok := <-c
	fmt.Printf("value: %v\nok: %v\n", v, ok)

	// pulls value 9 from c, returns true ok
	v, ok = <-c
	fmt.Printf("value: %v\nok: %v\n", v, ok)

	// pulls value 0 from c, returns false ok
	v, ok = <-c
	fmt.Printf("value: %v\nok: %v\n", v, ok)
}
