package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err occurred: ", err) // prints err with date/time stamp
	}
}
