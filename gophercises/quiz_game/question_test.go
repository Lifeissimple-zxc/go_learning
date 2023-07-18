package main

import (
	"testing"
)

var validQ = QuizQ{
	q: "5+5",
	a: "10",
}

// Tests valid slice of string parsing to a QuizQ struct
func TestParseQValid(t *testing.T) {
	inp := []string{"5+5", "10 "}
	res, _ := ParseQ(inp) // No error check as input is controlled

	if res != validQ {
		t.Errorf("Expected: %v. Got: %v", validQ, res)
	}

}

// Tests parsing invalid input
func TestParseQInvalid(t *testing.T) {
	inp := []string{"5+5", "10 ", "Unexpected value"}
	_, err := ParseQ(inp) // No error check as input is controlled

	if err == nil {
		t.Errorf("Error cannot be non-nil when parsing an input of non 2 len.")
	}

}

func TestPrintStr(t *testing.T) {
	exp := "5+5 | "
	res := validQ.PrintStr()

	if res != exp {
		t.Errorf("Wrong output for QuizQ PrintStr()")
	}

}

// Tests correct answer incrementing points by 1
func TestCheckAnswerCorrect(t *testing.T) {
	points := 1
	usrA := "10"
	validQ.CheckAnswer(&points, usrA)
	if points != 2 {
		t.Errorf("CheckAnswer did not account for correct answer correctly. Points: %d", points)
	}
}

// Tests incorrect answer leaving points unchanges
func TestCheckAnswerIncorrect(t *testing.T) {
	points := 1
	usrA := "blabla some wrong answer"
	validQ.CheckAnswer(&points, usrA)
	if points != 1 {
		t.Errorf("CheckAnswer did not account for correct answer correctly. Points: %d", points)
	}
}
