package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrMessage = errors.New("Can't take square root of negative number!")

func main() {
	fmt.Printf("\n%T\n", ErrMessage)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMessage
	}
	return 42, nil
}
