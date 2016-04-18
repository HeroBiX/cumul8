package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Suit   string `json:"suit"`
	Number int    `json:"number"`
}

//deck
var cards = []Card{{Suit: "Spades", Number: 2}, {Suit: "Spades", Number: 3}, {Suit: "Spades", Number: 4}, {Suit: "Spades", Number: 5}, {Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}, {Suit: "Spades", Number: 14}, {Suit: "Hearts", Number: 2}, {Suit: "Hearts", Number: 3}, {Suit: "Hearts", Number: 4}, {Suit: "Hearts", Number: 5}, {Suit: "Hearts", Number: 6}, {Suit: "Hearts", Number: 7}, {Suit: "Hearts", Number: 8}, {Suit: "Hearts", Number: 9}, {Suit: "Hearts", Number: 10}, {Suit: "Hearts", Number: 11}, {Suit: "Hearts", Number: 12}, {Suit: "Hearts", Number: 13}, {Suit: "Hearts", Number: 14}, {Suit: "Diamonds", Number: 2}, {Suit: "Diamonds", Number: 3}, {Suit: "Diamonds", Number: 4}, {Suit: "Diamonds", Number: 5}, {Suit: "Diamonds", Number: 6}, {Suit: "Diamonds", Number: 7}, {Suit: "Diamonds", Number: 8}, {Suit: "Diamonds", Number: 9}, {Suit: "Diamonds", Number: 10}, {Suit: "Diamonds", Number: 11}, {Suit: "Diamonds", Number: 12}, {Suit: "Diamonds", Number: 13}, {Suit: "Diamonds", Number: 14}, {Suit: "Clubs", Number: 2}, {Suit: "Clubs", Number: 3}, {Suit: "Clubs", Number: 4}, {Suit: "Clubs", Number: 5}, {Suit: "Clubs", Number: 6}, {Suit: "Clubs", Number: 7}, {Suit: "Clubs", Number: 8}, {Suit: "Clubs", Number: 9}, {Suit: "Clubs", Number: 10}, {Suit: "Clubs", Number: 11}, {Suit: "Clubs", Number: 12}, {Suit: "Clubs", Number: 13}, {Suit: "Clubs", Number: 14}}

// shuffle 1 deck and then shuffle a 6 deck shoe
// 6 deck shoe is based on this YT video: https://www.youtube.com/watch?v=tpv5sqoveuc
func main() {
	fmt.Println("Shuffle 1 deck of Cards")
	Shuffle(cards)

	fmt.Println("Time to shuffle a 6 deck shoe")
	deckShoe(cards)
}

// 6 Deck Shoe
func deckShoe(d []Card) {
	d = creatingSixDeckShoe(d) // adding 5 decks of cards

	deck1, deck2 := splitDeck(d) // Splitting the deck into 2
	mergeDecks(deck1, deck2)

}

// adding so the deck of card contains 6 packs of cards (total of 312 cards)
func creatingSixDeckShoe(d []Card) []Card {
	for i := 0; i <= 4; i++ {
		cards = append(cards, d...)
	}
	fmt.Println(len(cards))
	return cards
}

// Split deck
func splitDeck(d []Card) ([]Card, []Card) {
	var deck1 []Card
	var deck2 []Card

	deck1 = append(deck1, d[:len(d)/2]...) // give half the deck to deck1
	deck2 = append(deck2, d[len(d)/2:]...) // give other half to deck2
	return deck1, deck2
}

//shuffle
func Shuffle(slc []Card) {
	for i := 1; i < len(slc); i++ {
		// r := rand.Seed(time.Now().UnixNano())
		r := rand.Intn(i + 1)
		fmt.Println("i: ", i)
		fmt.Println("r: ", r)
		fmt.Println("slc[r]: ", slc[r])
		fmt.Println("slc[i]", slc[i])
		if i != r {
			fmt.Println("!=")
			slc[r], slc[i] = slc[i], slc[r]
			fmt.Println("slc[r]: ", slc[r])
			fmt.Println("slc[i]", slc[i])
		}
		fmt.Println("slc :", slc)
	}
	fmt.Println(len(slc))
}

func mergeDecks(deck1, deck2 []Card) {
	fmt.Println("Deck1 len", len(deck1), "Deck2 len", len(deck2))

}

// riffle the cards
func riffle(unRiffle []Card) []Card {
	var riffled []Card     // new deck with riffled cards
	q := len(unRiffle) / 4 // how many cards is a quarter

	for i := 0; i <= 3; i++ {
		riffled = append(riffled, unRiffle[len(unRiffle)-q:]...)                   // taking a quarter of the top cards from unRiffle and put them into riffled
		unRiffle = append(unRiffle[:len(unRiffle)-q], unRiffle[len(unRiffle):]...) // remove those cards from unRiffle
	}
	return riffled // return the riffled deck of cards
}
