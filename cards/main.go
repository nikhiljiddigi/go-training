package main

func main() {
	cards := newDeck()
	// hand, _ := deal(cards, 6)
	// cards.saveToFile("myDeck.txt")
	// myDeck := newDeckFromFile("myDefck.txt")
	// myDeck.print()
	// remainingCards.toString()
	cards.shuffle()
	cards.print()
}
