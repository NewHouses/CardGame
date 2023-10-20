package main

func main() {
	var cards deck

	chatBot := botFactory()

	if chatBot.askCreateNewDeck() {
		cards = createAndSaveNewDeck()
	} else {
		cards = createDeckFromFile(chatBot)
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

func createDeckFromFile(chatBot bot) deck {
	filepath := chatBot.askFilePath()
	return newDeckFromFile(filepath)
}
