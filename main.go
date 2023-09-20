package main

func main() {

	card := "Ace of Spades"

	cards := deck{card, newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
