package main

func main() {
	cards := newDeck()

	cards.print()

	hand, remamingDeck := deal(cards, 5)

	hand.print()
	remamingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
