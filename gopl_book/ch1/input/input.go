package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // to convert strings to numbers
	"time"
)

func main() {
	// Create scanner for reading user input from CLI interface
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your year of birth: ")
	scanner.Scan() // actually scan for input
	// scanner always returns a string, even 3.14 will be a string
	// Store our data to a variable & convert to an int
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	// Error check
	if err != nil {
		fmt.Printf("Error when parsing input: %v\n", err)
		os.Exit(1)
	}
	// '%q' stands for single quoted character literal escaped with GO syntax
	// fmt.Printf("You typed %q year.\n", input)
	// Compute age (simplified)
	yr := time.Now().Year()
	fmt.Printf(
		"Current year: %d. Your age: %d.\n",
		yr, yr-int(input),
	)

}
