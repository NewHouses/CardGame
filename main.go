package main

import (
	"fmt"
	"os"
)

func main() {
	var answer string
	var cards deck
	var chatBot bot
	fmt.Println("English/Galego? (en/gal)")
	fmt.Scanln(&answer)
	if answer == "en" {
		chatBot = englishBot{}
	} else if answer == "gal" {
		chatBot = galicianBot{}
	} else {
		fmt.Println("Answer " + answer + " is not valid / Resposta " + answer + " non válida")
		os.Exit(1)
	}

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
