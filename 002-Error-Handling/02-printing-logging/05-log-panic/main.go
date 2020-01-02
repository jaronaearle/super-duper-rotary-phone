package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txr")
	if err != nil {
		log.Panic(err)
	}
}
