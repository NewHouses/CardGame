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
		cards = createAndSaveNewDeck()
	} else if answer == "n" {
		cards = createDeckFromFile()
	} else {
		fmt.Println("Not valid answer")
		os.Exit(1)
	}

	cards.shuffle()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("My hand")
	hand.print()
	fmt.Println("\n\nRemaining deck")
	remainingDeck.print()
}

func createAndSaveNewDeck() deck {
	cards := newDeck()
	cards.saveToFile("my_cards")
	return cards
}

func createDeckFromFile() deck {
	var filename string
	fmt.Println("Introduce the deck file path?")
	fmt.Scanln(&filename)
	return newDeckFromFile(filename)
}
