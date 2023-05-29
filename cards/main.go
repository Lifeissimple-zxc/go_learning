// This means it is executable, not reausable
package main

const CARDS_FILE_NAME string = "cards_deck.txt"

// Needed for main package for code to be executed
func main() {
	// Read from a file
	// cards := newDeckFromFile(CARDS_FILE_NAME)
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// fmt.Println("Hand is")
	// hand.print()

	// fmt.Println("Remaining is")
	// remainingDeck.print()

	// fmt.Println("Saving cards to file")
	// cards.saveToFile(CARDS_FILE_NAME)

}
