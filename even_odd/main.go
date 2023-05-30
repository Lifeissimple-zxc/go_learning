package main

import "fmt"

const END int = 11

func main() {
	i := 0
	// while loop pattern
	for i < END {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
		// Increment our i to check a different number every time
		i++
	}
}
