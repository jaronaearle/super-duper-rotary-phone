package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	defer func() {
		fmt.Println("The call to os.Exit() cancels this function call")
	}()

	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	// the Fatalln is equal to calling Println() followed by calling os.Exit(1)
}
