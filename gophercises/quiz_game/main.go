package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"
)

const (
	startPr    = "Press ENTER to start the quiz\n."
	enterTries = 20
	timeoutMsg = "\nQuiz time ended :(. Results:"
	intrptMsg  = "\nQuiz was interrupted. Results:"
)

func main() {

	// Profiling
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// Parse args: flags package
	csvPtr := flag.String("csv", "quiz.csv", "Path to CSV of q (question), a (answer) schema.")
	timer := flag.Int("timer", 30, "Time limit for 1 question (seconds).") // this is for part 2!
	flag.Parse()

	// Init variables we need to give a summary on the quiz to users
	var points int // Here we will track # of correct answers!
	var qsCnt int  // Here we will track # of questions in a quiz

	// Read CSV: CSV package
	f, err := os.Open(*csvPtr)
	if err != nil {
		fmt.Printf("Failed to open %s. Details: %v\n", *csvPtr, err)
		os.Exit(1)
	}
	defer f.Close()
	// Get total # of questions
	qs := readQuizFromCsv(f)
	qsCnt = len(qs) - 1 // Deduct header row bc it is not a question!
	fmt.Printf("Found %d questions in %s. Time to complete: %d!\n", qsCnt, *csvPtr, *timer)

	promptStart()

	// Set timer
	timeout := time.NewTimer(time.Duration(*timer) * time.Second)
	// Set interrupt signal
	intrpt := make(chan os.Signal, 1)
	signal.Notify(intrpt, os.Interrupt)

	// Schedule signals to stop quiz
	go stopQuiz(&points, &qsCnt, *timeout, intrpt)
	// Ask questions to the user
	askQs(qs, &points, &qsCnt)

	fmt.Println(getSummary(points, qsCnt))

}

// Prompts use to press ENTER to start the quiz.
// Terminates after n wrong inputs (set in const)
func promptStart() {
	var startS string // Here we will save user input for starting the quiz
	for i := 0; i < enterTries; i++ {
		fmt.Printf(startPr)
		fmt.Scanln(&startS)
		if startS == "" {
			break
		}
	}
}

// Iteratively asks questions using QuizQ struct
func askQs(qs [][]string, pointsPtr, qsCntPtr *int) {
	// Iterate over rows, we skip 0 bc it contains header
	for _, row := range qs[1:] {
		q, err := ParseQ(row)
		if err != nil {
			fmt.Printf("Failed to pause question from row: %v", row)
			*qsCntPtr--
			continue
		}
		q.Ask(pointsPtr)
	}
}

func stopQuiz(pointsPtr, qsCntPtr *int, t time.Timer, intrpt chan os.Signal) {
	select {
	// User ran out of time
	case <-t.C:
		fmt.Println(timeoutMsg)
		fmt.Println(getSummary(*pointsPtr, *qsCntPtr))
		os.Exit(0)
	// Manual interrupt signal received
	case <-intrpt:
		fmt.Println(intrptMsg)
		fmt.Println(getSummary(*pointsPtr, *qsCntPtr))
		os.Exit(0)
	}
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
