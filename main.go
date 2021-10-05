package main

func main() {
	cards := newCard()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	cards.saveToFile("myfile")
	new_cards := newDeckFromFile("myfile")
	new_cards.shuffle()
	new_cards.print()

}
