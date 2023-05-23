package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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
