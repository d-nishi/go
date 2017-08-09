package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Type string
	Suit string
}

type Deck []Card

//create a new deck
func New() (deck Deck) {
	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	suits := []string{"Heart", "Spade", "Diamond", "Club"}

	for i := 0; i < len(types); i++ {
		for j := 0; j < len(suits); j++ {
			card := Card{
				Type: types[i],
				Suit: suits[j],
			}
			deck = append(deck, card)
		}
	}
	return
}

//shuffle the deck
func Shuffle(deck Deck) Deck {
	for i := 0; i < len(deck); i++ {
		r := rand.Intn(i + 1)

		if i != r {
			deck[i], deck[r] = deck[r], deck[i]
		}
	}
	return deck
}

//display deck contents
func Show(deck Deck) Deck {
	for i := 0; i < len(deck); i++ {
		fmt.Println(deck[i])
	}
	return deck
}

func main() {

	deck := New()
	Show(Shuffle(deck))
}
