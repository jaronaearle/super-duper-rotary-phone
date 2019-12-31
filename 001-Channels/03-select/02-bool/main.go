package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
		close(q)
	}
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("value received from even channel: ", v)
		case v := <-o:
			fmt.Println("value received from odd channel: ", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok", i, ok)
				return
			} else {
				fmt.Println("from comma ok: ", i)
			}
		}
	}
}
