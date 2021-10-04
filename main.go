package main

func main() {
	cards := newCard()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
