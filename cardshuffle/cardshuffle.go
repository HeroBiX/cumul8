package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit   string `json:"suit"`
	Number int    `json:"number"`
}

//deck
var cards = []Card{{Suit: "Spades", Number: 2}, {Suit: "Spades", Number: 3}, {Suit: "Spades", Number: 4}, {Suit: "Spades", Number: 5}, {Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}, {Suit: "Spades", Number: 14}, {Suit: "Hearts", Number: 2}, {Suit: "Hearts", Number: 3}, {Suit: "Hearts", Number: 4}, {Suit: "Hearts", Number: 5}, {Suit: "Hearts", Number: 6}, {Suit: "Hearts", Number: 7}, {Suit: "Hearts", Number: 8}, {Suit: "Hearts", Number: 9}, {Suit: "Hearts", Number: 10}, {Suit: "Hearts", Number: 11}, {Suit: "Hearts", Number: 12}, {Suit: "Hearts", Number: 13}, {Suit: "Hearts", Number: 14}, {Suit: "Diamonds", Number: 2}, {Suit: "Diamonds", Number: 3}, {Suit: "Diamonds", Number: 4}, {Suit: "Diamonds", Number: 5}, {Suit: "Diamonds", Number: 6}, {Suit: "Diamonds", Number: 7}, {Suit: "Diamonds", Number: 8}, {Suit: "Diamonds", Number: 9}, {Suit: "Diamonds", Number: 10}, {Suit: "Diamonds", Number: 11}, {Suit: "Diamonds", Number: 12}, {Suit: "Diamonds", Number: 13}, {Suit: "Diamonds", Number: 14}, {Suit: "Clubs", Number: 2}, {Suit: "Clubs", Number: 3}, {Suit: "Clubs", Number: 4}, {Suit: "Clubs", Number: 5}, {Suit: "Clubs", Number: 6}, {Suit: "Clubs", Number: 7}, {Suit: "Clubs", Number: 8}, {Suit: "Clubs", Number: 9}, {Suit: "Clubs", Number: 10}, {Suit: "Clubs", Number: 11}, {Suit: "Clubs", Number: 12}, {Suit: "Clubs", Number: 13}, {Suit: "Clubs", Number: 14}}

// adding so the deck of card contains 6 packs of cards (total of 312 cards)
func creatingSixDeckShoe(d []Card) []Card {
	for i := 0; i <= 4; i++ {
		cards = append(cards, d...)
	}
	fmt.Println(len(cards))
	return cards
}

// Split deck in two
func SplitDeck(d []Card) ([]Card, []Card) {
	var deck1 []Card
	var deck2 []Card

	deck1 = append(deck1, d[:len(d)/2]...) // give half the deck to deck1
	deck2 = append(deck2, d[len(d)/2:]...) // give other half to deck2
	fmt.Println("Deck 1 ", len(deck1), "Deck2 ", len(deck2))
	return deck1, deck2
}

// Take out 1 quarter of the deck
func QuarterDeck(q int, deck1, deck2 []Card) ([]Card, []Card, []Card) {
	var tempDeck []Card // new deck with riffled cards

	tempDeck = append(tempDeck, deck1[len(deck1)-q:]...)        // move cards from deck1 to tempDeck
	deck1 = append(deck1[:len(deck1)-q], deck1[len(deck1):]...) // remove those cards from deck1
	tempDeck = append(tempDeck, deck2[len(deck2)-q:]...)        // move cards from deck2 to tempDeck
	deck2 = append(deck2[:len(deck2)-q], deck2[len(deck2):]...) // remove those cards from deck2

	return deck1, deck2, tempDeck
}

//shuffle
func Shuffle(slc []Card) []Card {

	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
	return slc
}

// Riffle
func Riffle(unRiffle []Card) []Card {
	var riffled []Card     // new deck with riffled cards
	q := len(unRiffle) / 4 // how many cards is a quarter

	for i := 0; i <= 3; i++ {
		riffled = append(riffled, unRiffle[len(unRiffle)-q:]...)                   // moving cards from unRiffle to Riffle
		unRiffle = append(unRiffle[:len(unRiffle)-q], unRiffle[len(unRiffle):]...) // remove those cards from unRiffle
	}
	riffled = append(riffled, unRiffle[:]...) // adding the left over cards into riffled
	return riffled                            // return the riffled deck of cards
}

// Strip Shuffle Strip
func ShuffleRiffleShuffle(tempDeck []Card) []Card {

	tempDeck = Shuffle(tempDeck)
	tempDeck = Riffle(tempDeck)
	tempDeck = Shuffle(tempDeck)
	return tempDeck
}

func CutDeck(deck []Card) []Card {
	var tempDeck []Card
	cutNumber := random(1, len(deck))

	tempDeck = append(tempDeck, deck[:cutNumber]...)
	deck = append(deck[:0], deck[cutNumber:]...) // delete those cards
	deck = append(deck[:], tempDeck[:]...)

	return deck
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// shuffle 1 deck and then shuffle a 6 deck shoe
// 6 deck shoe is based on this YT video: https://www.youtube.com/watch?v=tpv5sqoveuc
func main() {
	rand.Seed(time.Now().Unix())
	var tempDeck []Card

	fmt.Println("Shuffle 1 deck of Cards")
	Shuffle(cards)

	fmt.Println("Time to create 6 deck shoe")
	tempDeck = creatingSixDeckShoe(cards) // adding 5 decks of cards
	deck1, deck2 := SplitDeck(tempDeck)   // Splitting the deck into 2
	tempDeck = shuffleDeckShoe(deck1, deck2)

	fmt.Println("Shuffled 6 deck shoe completed")
	fmt.Println(tempDeck)
	fmt.Println(len(tempDeck))
}

// Shuffle the 6 Deck Shoe
func shuffleDeckShoe(deckLeft, deckRight []Card) []Card {
	fmt.Println("deckLeft len", len(deckLeft), "deckRight len", len(deckRight))
	q := len(deckLeft) / 4 // how many cards is a quarter
	var workingDeck []Card
	var tempDeck []Card

	deckLeft, deckRight, tempDeck = QuarterDeck(q, deckLeft, deckRight) // Take a quarter from each pile
	workingDeck = append(workingDeck, tempDeck[:]...)                   // merge those together to the working deck
	workingDeck = ShuffleRiffleShuffle(workingDeck)

	deckRight, deckLeft = Combine2Decks(deckLeft, deckRight, workingDeck, q)
	workingDeck = RiffleTogetherDecks(deckLeft, deckRight, q)

	workingDeck = CutDeck(workingDeck) // Cut the deck
	tempDeck = append(tempDeck, workingDeck[:1]...)
	workingDeck = append(workingDeck[:0], workingDeck[1:]...) // discard the first card

	return workingDeck // Return to main

}

func Combine2Decks(deckLeft, deckRight, workingDeck []Card, q int) ([]Card, []Card) {
	var tempDeck []Card
	var isRightDeck bool // should the cards be taken from the right or left stack

	for i := 0; i <= 5; i++ { // loop through the cards and shuffle them into the working deck
		if isRightDeck == true { // take 1/4 from working deck and right deck
			workingDeck, deckRight, tempDeck = QuarterDeck(q, workingDeck, deckRight)
			isRightDeck = false

		} else { // take 1/4 from working deck and left deck

			workingDeck, deckLeft, tempDeck = QuarterDeck(q, workingDeck, deckLeft)
			isRightDeck = true
		}

		tempDeck = ShuffleRiffleShuffle(tempDeck)         // shuffle, Riffle, shuffle
		workingDeck = append(workingDeck, tempDeck[:]...) // add tempDeck to workingDeck
		// workingDeck = Riffle(workingDeck)                 // Riffle working deck
		tempDeck = tempDeck[:0] // clean out the tempDeck

	}

	deckRight, deckLeft = SplitDeck(workingDeck)
	workingDeck = workingDeck[:0] // work deck becomes empty

	return deckRight, deckLeft

}

func RiffleTogetherDecks(deckRight, deckLeft []Card, q int) []Card {
	var tempDeck []Card
	var workingDeck []Card

	for i := 0; i <= 3; i++ { // Riffle together left and right deck

		deckRight, deckLeft, tempDeck = QuarterDeck(q, deckRight, deckLeft)
		tempDeck = Riffle(tempDeck)
		workingDeck = append(workingDeck, tempDeck[:]...)
		tempDeck = tempDeck[:0] // clean out the tempDeck
	}
	return workingDeck
}
