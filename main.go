package main

import (
	"fmt"
	"os"
)

func main() {
	var answer string
	var cards deck
	fmt.Println("Do you want to create a new Deck? (y/n)")
	fmt.Scanln(&answer)

	if answer == "y" {
		cards = newDeck()
		cards.saveToFile("my_cards")
	} else if answer == "n" {
		var filename string
		fmt.Println("Introduce the deck file path?")
		fmt.Scanln(&filename)
		cards = newDeckFromFile(filename)
	} else {
		fmt.Println("Not valid answer")
		os.Exit(1)
	}

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("My hand")
	hand.print()
	fmt.Println("\n\nRemaining deck")
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
