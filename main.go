package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("My hand")
	hand.print()
	fmt.Println("\n\nRemaining deck")
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
