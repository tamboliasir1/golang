package main

import "fmt"

//create a new type of deck which is slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardsuits := []string{"Daimonds", "Spades", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "two", "three", "four"}

	for _, suits := range cardsuits {
		for _, values := range cardValues {
			cards = append(cards, suits+" of "+values)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
