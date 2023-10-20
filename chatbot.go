package main

import (
	"fmt"
)

type bot interface {
	askCreateDeck(pointerAnswer *string)
	showDecks(hand deck, remainingDeck deck)
}

type englishBot struct{}

func (englishBot) askCreateDeck(pointerAnswer *string) {
	fmt.Println("Do you want to create a new Deck? (y/n)")
	fmt.Scanln(&*pointerAnswer)
}

func (englishBot) showDecks(hand deck, remainingDeck deck) {
	fmt.Println("My hand")
	hand.print()
	fmt.Println("\n\nRemaining deck")
	remainingDeck.print()
}

type galicianBot struct{}

func (galicianBot) askCreateDeck(pointerAnswer *string) {
	fmt.Println("Queres crear unha nova baralla? (y/n)")
	fmt.Scanln(&*pointerAnswer)
}

func (galicianBot) showDecks(hand deck, remainingDeck deck) {
	fmt.Println("A miña mao")
	hand.print()
	fmt.Println("\n\nResto da baralla")
	remainingDeck.print()
}
