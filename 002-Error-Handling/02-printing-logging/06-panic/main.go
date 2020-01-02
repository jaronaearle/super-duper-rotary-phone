package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txr")
	if err != nil {
		panic(err)
	}
}
