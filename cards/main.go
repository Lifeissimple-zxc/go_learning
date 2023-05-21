// This means it is executable, not reausable
package main

// Needed for main package for code to be executed
func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// Iterate over cards
	cards.print() // Looks like OOP?
}

func newCard() string {
	return "Five of Diamonds"
}
