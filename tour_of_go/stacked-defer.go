package main

import "fmt"

func main() {
	fmt.Println("Starting count!")
	// Many defers
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	// The defers are pushed to a stack to we get results back in a LIFO order!
	fmt.Println("Done!")
}
