package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four",} 
	// "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{
	
	return strings.Join([]string(d),",")
}

func (d deck ) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0600)
	
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)	
		os.Exit(1)
	}

	s := strings.Split(string(bs),",")


	return deck(s)
}