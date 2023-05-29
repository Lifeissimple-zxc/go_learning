package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Creating a new custom deck type (it is a slice of strings)
type deck []string // It kinda extends the behavior of string Slice

func newDeck() deck {
	cards := deck{}

	// Inputs for our combinations
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Iterative generation of cards
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)

		}
	}

	return cards
}

func (d deck) print() { // (d deck) means receiver of type deck
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Return multiple values from the function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Type conversion & Join string slice to a string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(
		filename,
		[]byte(d.toString()),
		0666,
	)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// Check for error
	if err != nil {
		// Option 1 - log error + return a new deck
		// Option 2 - no new deck ==> terrible error (log it) ==> quit the program
		// The author goes for option 2 in the course
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// Getting here means we can return our dec
	// We need to convert bytes to a string before we return!
	s := strings.Split(string(bs), ",")
	// Type conversion from a slice of string to a deck
	return deck(s)
}

func (d deck) shuffle() {
	// Precautions to make our shuffling truly random
	// Create our source (seed) using unix nanoseconds as input
	// Using unix nano makes the source truly random
	source := rand.NewSource(time.Now().UnixNano())
	// Create our new random generator
	r := rand.New(source)
	// Modifies order of cards in the deck d
	for i := range d {
		// Generate a random integer
		// newPosition := rand.Intn(len(d) - 1) // not truly random
		newPosition := r.Intn(len(d) - 1)
		// Swap card posititons
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
