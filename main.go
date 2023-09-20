package main

import "fmt"

func main() {

	card := "Ace of Spades"

	cards := []string{card, newCard()}
	cards = append(cards, getState())

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(card)
	printState()
}

func newCard() string {
	return "Five of Diamonds"
}
