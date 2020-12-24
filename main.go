package main

func main() {
	cards := newDeck()
	cards.addCard(newCard("red"))
	cards.addCard("King")

	hand, restDeck := cards.deal(4)

	hand.print()
	restDeck.print()

	// cards.print()
}

func newCard(color string) string {
	return "Ace " + color
}
