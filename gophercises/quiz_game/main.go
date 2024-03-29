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
	"strings"
	"time"
)

const (
	startPr    = "Press ENTER to start the quiz\n."
	enterTries = 20
	timeoutMsg = "\nQuiz time ended. Results:"
	intrptMsg  = "\nQuiz was interrupted. Results:"
)

func main() {

	// Profiling
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// Parse args: flags package
	csvPtr := flag.String("csv", "quiz.csv", "Path to CSV of q (question), a (answer) schema.")
	timer := flag.Int("timer", 30, "Time limit for the quiz (seconds).")
	flag.Parse()

	// Init variables we need to give a summary on the quiz to users
	var points int // Here we will track # of correct answers!
	var qsCnt int  // Here we will track # of questions in a quiz

	// Read CSV: CSV package
	f, err := os.Open(*csvPtr)
	if err != nil {
		errExit(fmt.Sprintf("Failed to open %s. Details: %v\n", *csvPtr, err))
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

// Prints msg to the user and calls os.Exit(1).
// Maily aimed at reducing repetitive code.
func errExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Prompts use to press ENTER to start the quiz.
// Terminates after n wrong inputs (set in const)
func promptStart() {
	var startS string // Here we will save user input for starting the quiz
	for i := 0; i < enterTries; i++ {
		fmt.Printf(startPr)
		// Parse text accounting for trailing whitespaces
		fmt.Scanf("%s\n", &startS)
		startS = strings.TrimSpace(startS)
		if startS == "" {
			break
		}
		// Default to an empty string to avoid a broken cycle
		if i < (enterTries - 1) {
			startS = ""
		} else {
			// Getting here means user did not comply with the input policy
			errExit(fmt.Sprintf("You did start the quiz. Last input: %s", startS))
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
		fmt.Printf("%s", q.PrintStr()) // Prints question to the user
		// Read answer from STDIN
		var usrA string
		fmt.Scanf("%s\n", &usrA)
		// Check if correct
		q.CheckAnswer(pointsPtr, usrA)
	}
}

// Enforces exit from the program on timeout or interrupt signal.
func stopQuiz(pointsPtr, qsCntPtr *int, t time.Timer, intrpt chan os.Signal) {
	select {
	// User ran out of time
	// <-t.C blocks until we get the message, that's why we place it on a separate GORO
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

// Reads questions from a csv to a nested slice of strings.
func readQuizFromCsv(r io.Reader) [][]string {
	// Get total # of questions
	csvR := csv.NewReader(r)
	qs, err := csvR.ReadAll()
	if err != nil {
		errExit(fmt.Sprintf("Failed to read data from %s. Details: %v\n", r, err))
	}
	return qs
}

// Prepares quiz summary string to be printed to the user.
func getSummary(points, qsCnt int) string {
	sucRate := math.Round(float64(points)/float64(qsCnt)*100) / 100
	sumStr := fmt.Sprintf("%d out of %d correct! %2.f%% success rate.", points, qsCnt, sucRate*100)
	return sumStr
}
