package main

import (
	"os"
	"testing"
)

const TEST_FILE_NAME string = "_cardDeckTest.txt"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	// Check for deck len
	dLen := len(d)
	if dLen != 16 {
		t.Errorf("Expected deck of 16, but got %v", dLen)
	}

	// Check individual cards in the deck
	// First card
	exp1 := "Ace of Spades"
	act1 := d[0]
	if act1 != exp1 {
		t.Errorf("Expected first card %v, but got %v", exp1, act1)
	}
	// Last card
	expL := "Four of Clubs"
	actL := d[dLen-1]
	if actL != expL {
		t.Errorf("Expected last card %v, but got %v", expL, actL)
	}
}

// File IO tests
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Delete file
	os.Remove(TEST_FILE_NAME)
	// Get New deck
	d := newDeck()
	// Write to file
	d.saveToFile(TEST_FILE_NAME)
	// Load from file
	loadedDeck := newDeckFromFile(TEST_FILE_NAME)
	// Assert
	lenLoaded, exp := len(loadedDeck), 16
	if lenLoaded != exp {
		t.Errorf("Expected len: %v, got len %v", exp, lenLoaded)
	}
	// Delete File
	os.Remove(TEST_FILE_NAME)
}
