package main

import (
	"errors"
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
		return QuizQ{}, errors.New(
			fmt.Sprintf("Wrong len of input: %d. Needs to be 2!", len(row)),
		)
	}
	newQ := QuizQ{
		q: row[0],
		a: strings.TrimSpace(strings.ToLower(row[1])),
	}
	return newQ, nil
}

// Converts question to a formatted string that can be printed to a user
func (q QuizQ) PrintStr() string {
	return fmt.Sprintf("%s | ", q.q)
}

// Removes leading and trailing spaces from usrA, converts it to lowercase.
// Then compares with q.a and increments value of pointsPtr by 1 if the two match.
func (q QuizQ) CheckAnswer(pointsPtr *int, usrA string) {
	usrA = strings.ToLower(strings.TrimSpace(usrA))
	if q.a == usrA {
		*pointsPtr++
	}
}
