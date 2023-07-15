package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
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
	qs := readQuizFromCsv(f)
	qsCnt := len(qs) - 1 // We need it later too
	fmt.Printf("Found %d questions in %s. Starting quiz with %d secods per question!\n", qsCnt, *csvPtr, *timer)

	var points int // Here we will track of correct answers!
	// Iterate over rows, we skip 0 bc it contains header
	for _, row := range qs[1:] {
		// TODO make row a separate datastrcutre with some receiver functions later
		q, err := ParseQ(row)
		if err != nil {
			fmt.Printf("Failed to pause question from row: %v", row)
			qsCnt--
			continue
		}
		q.Ask(&points)
	}

	fmt.Println(getSummary(points, qsCnt))
}

// Reads questions from a csv to a nested slice of strings
func readQuizFromCsv(r io.Reader) [][]string {
	// Get total # of questions
	csvR := csv.NewReader(r)
	qs, err := csvR.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read data from %s. Details: %v\n", r, err)
		os.Exit(1)
	}
	return qs
}

// Prepares quiz summary string to be printed to the user
func getSummary(points, qsCnt int) string {
	sucRate := math.Round(float64(points)/float64(qsCnt)*100) / 100
	sumStr := fmt.Sprintf("%d out of %d correct! %2.f success rate.", points, qsCnt, sucRate*100)
	return sumStr
}
