package main

import "fmt"

func main() {
	f()
	fmt.Println("returned normally from fun f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling func g")
	g(0)
	fmt.Println("Returned normally from func g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("panicking! panicking! panicking!")
		panic(fmt.Sprintf("%v\n", i))
	}
	defer fmt.Println("defer in func g", i)
	fmt.Println("printing in func g", i)
	g(i + 1)
}
