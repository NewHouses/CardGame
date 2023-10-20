package main

import (
	"fmt"
	"os"
)

type bot interface {
	askCreateNewDeck() bool
	notValidAnswer(answer string)
	askFilePath() string
	showDecks(hand deck, remainingDeck deck)
}

func botFactory() bot {
	var answer string
	fmt.Println("English/Galego? (en/gal)")
	fmt.Scanln(&answer)

	if answer != "en" && answer != "gal" {
		fmt.Println("Answer " + answer + " is not valid / Resposta " + answer + " non válida")
		os.Exit(1)
	} else if answer == "gal" {
		return galicianBot{}
	}

	return englishBot{}
}

type englishBot struct{}

func (eb englishBot) askCreateNewDeck() bool {
	var answer string
	fmt.Println("Do you want to create a new Deck? (y/n)")
	fmt.Scanln(&answer)

	if answer != "y" && answer != "n" {
		eb.notValidAnswer(answer)
	} else if answer == "n" {
		return false
	}

	return true
}

func (englishBot) notValidAnswer(answer string) {
	fmt.Println("Answer " + answer + " is not valid")
	os.Exit(1)
}

func (englishBot) askFilePath() string {
	var filename string
	fmt.Println("Introduce the deck file path")
	fmt.Scanln(&filename)

	return filename
}

func (englishBot) showDecks(hand deck, remainingDeck deck) {
	fmt.Println("My hand")
	hand.print()
	fmt.Println("\n\nRemaining deck")
	remainingDeck.print()
}

type galicianBot struct{}

func (gb galicianBot) askCreateNewDeck() bool {
	var answer string
	fmt.Println("Queres crear unha nova baralla? (s/n)")
	fmt.Scanln(&answer)

	if answer != "s" && answer != "n" {
		gb.notValidAnswer(answer)
	} else if answer == "n" {
		return false
	}

	return true
}

func (galicianBot) notValidAnswer(answer string) {
	fmt.Println("Resposta " + answer + " non válida")
	os.Exit(1)
}

func (galicianBot) askFilePath() string {
	var filename string
	fmt.Println("Introduce a ruta do arquivo da baralla")
	fmt.Scanln(&filename)

	return filename
}

func (galicianBot) showDecks(hand deck, remainingDeck deck) {
	fmt.Println("A miña mao")
	hand.print()
	fmt.Println("\n\nResto da baralla")
	remainingDeck.print()
}
