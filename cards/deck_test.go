package main

import "testing"

func TestNewDeck(t *testing.T) {
	d = newDeck()
	// Check for deck len
	dLen := len(d)
	if dLen != 16 {
		t.Errorf("Expected deck of 16, but got", dLen)
	}

}
