package main

import (
	"fmt"
)

func main() {
	var answer string
	var cards deck

	chatBot := botFactory()

	chatBot.askCreateDeck(&answer)

	if answer == "y" {
		cards = createAndSaveNewDeck()
	} else if answer == "n" {
		cards = createDeckFromFile()
	} else {
		chatBot.notValidAnswer(answer)
	}

	cards.shuffle()

	hand, remainingDeck := deal(cards, 5)

	chatBot.showDecks(hand, remainingDeck)
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
