package main

import "fmt"

// Creating a new custom deck type (it is a slice of strings)
type deck []string // It kinda extends the behavior of string Slice

func (d deck) print() { // (d deck) means receiver of type deck
	for i, card := range d {
		fmt.Println(i, card)
	}
}
