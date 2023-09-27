package main

// import "fmt"

func main() {

	// cards := newDeck()
	// cards.saveToFile("cardDeck")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	cards := newDeckFromFile("cardDeck")
	cards.print()
}
