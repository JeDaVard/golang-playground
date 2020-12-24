package main

func main() {
	deck := newDeck()
	str, de := deck.toString()
	deckk := de.toArr(str)

	ioOperation.save("sdfgh", deckk)

}

func newCard(color string) string {
	return "Ace " + color
}
