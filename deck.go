package main

import "fmt"

type deck []string

func newCard() deck {
	cards := deck{}
	cardTypes := []string{"Clubs", "Spade", "Diamonds", "Hearts"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four"}

	for _, cardType := range cardTypes {
		for _, number := range cardNumbers {
			cards = append(cards, cardType+" of "+number)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}
