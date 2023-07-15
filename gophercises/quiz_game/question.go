package main

import (
	"fmt"
	"strings"
)

type QuizQ struct {
	q string
	a string
}

// Parses csv row to q quiz struct
func ParseQ(row []string) (QuizQ, error) {
	// Naive check for input's format
	if len(row) != 2 {
		return QuizQ{}, nil
	}
	newQ := QuizQ{
		q: row[0],
		a: strings.ToLower(row[1]),
	}
	return newQ, nil
}

// Asks the question, validates the answers.
// If the answer is correct: increments pToPoints by 1
func (q QuizQ) Ask(pToPoints *int) {
	// Variable to collect user input
	var usrA string
	fmt.Printf("%s | ", q.q) // Print question to the user
	fmt.Scanln(&usrA)
	// Validate input & increment points
	if q.a == usrA {
		*pToPoints++
	}
}
