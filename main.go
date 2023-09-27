package main

import "fmt"

func main() {

	cards := newDeck()
	cards.saveToFile("cardDeck")
	cards = newDeckFromFile("cardDeck")

	cards.shuffle()
	cards.print()

	hand, _ := deal(cards, 5)


	fmt.Println("HAND")
	hand.print()
	// fmt.Println("REST OF PACK")
	// remainingCards.print()

	
}
