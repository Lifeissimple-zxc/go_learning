package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a container for counting strings with key as string and int as value
	counts := make(map[string]int)
	// Get input type
	input := bufio.NewScanner(os.Stdin)
	// Scan input
	for input.Scan() {
		// Increment count for the string by changing value of the key
		counts[input.Text()]++
		fmt.Println(counts)
	}

	// Print the map's summary
	for line, n := range counts {
		if n > 1 {
			/*
				Format Comment
				%d is a decimal integer
				%s is an uninterpreted bytes of the string or slice
			*/
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
