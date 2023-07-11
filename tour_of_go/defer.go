package main

import "fmt"

func main() {
	defer fmt.Println("world") // evaluated immediately, but the execution happens when main() returns!

	fmt.Println("Hello")
}
