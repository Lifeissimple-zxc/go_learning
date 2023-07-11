package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Parse args: flags package
	csvPtr := flag.String("csv", "quiz.csv", "Path to CSV of q (question), a (answer) schema.")
	timer := flag.Int("timer", 10, "Time limit for 1 question (seconds).") // this is for part 2!
	flag.Parse()
	// Read CSV: CSV package
	f, err := os.Open(*csvPtr)
	if err != nil {
		fmt.Printf("Failed to open %s. Details: %v\n", *csvPtr, err)
		os.Exit(1)
	}
	defer f.Close()
	// Get total # of questions
	csvR := csv.NewReader(f)
	qs, err := csvR.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read data from %s. Details: %v\n", *csvPtr, err)
		os.Exit(1)
	}
	qsCnt := len(qs) - 1 // We need it later too
	fmt.Printf("Found %d questions in %s. Starting quiz with %d secods per question!\n", qsCnt, *csvPtr, *timer)
	var points int // Here we will track of correct answers!
	// Iterate over rows
	// We skip index 0 bc it contains header

	for _, row := range qs[1:] {
		// TODO make row a separate datastrcutre with some receiver functions later
		// Like askQuestion
		// validateAnswer
		// Parse questions? Might be actually not irrelevant, just keep everything lowercase
		q := row[0]
		a := strings.ToLower(row[1])
		// Collect user input
		var usrA string
		fmt.Printf("%s | ", q)
		fmt.Scanln(&usrA)
		// Validate input
		if a == usrA {
			points++
		}
	}

	printSummary(points, qsCnt)

}

// Prints quiz result to the user
func printSummary(points, qsCnt int) {
	sucRate := float64(points) / float64(qsCnt)
	fmt.Printf("%d out of %d correct! %2.f success rate.", points, qsCnt, sucRate*100)
}
