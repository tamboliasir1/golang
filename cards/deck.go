package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of deck which is slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardsuits := []string{"Ace", "Two", "Three", "Four"}
	cardValues := []string{"Daimonds", "Spades", "Hearts", "Clubs"}

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
