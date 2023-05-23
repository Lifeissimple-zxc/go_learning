// This means it is executable, not reausable
package main

import "fmt"

const CARDS_FILE_NAME string = "cards_deck.txt"

// Needed for main package for code to be executed
func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("Hand is")
	hand.print()

	fmt.Println("Remaining is")
	remainingDeck.print()

	fmt.Println("Saving cards to file")
	cards.saveToFile(CARDS_FILE_NAME)
}
