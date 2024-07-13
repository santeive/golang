package main

import "fmt"

// Create a new type of deck
// Which is a slice of strings

// Defining a deck: Is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// This is a receiver
// Where d is kind of ".this" or ".self" in javascript or python
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
