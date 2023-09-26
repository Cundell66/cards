package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	card := "";

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			card = value + " of " + suit
			cards = append(cards, card)
		}
	}

	return cards
}

func (d deck) print(){ 
	for _, card := range d{
		fmt.Println(card)
	}
}