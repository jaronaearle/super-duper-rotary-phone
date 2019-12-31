package main

import "fmt"

func main() {
	// channels are reference type, references underlying data
	c := make(chan int)
	// send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		// if left unclosed, the range loop below will try to pull values from the channel that do not exists and creates a deadlock
		close(c)
	}()

	//receive
	// range loop pulls values from the channel until then channel is closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")

}
