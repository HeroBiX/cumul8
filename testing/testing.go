package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Suit   string `json:"suit"`
	Number int    `json:"number"`
}

var cards1 = []Card{{Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}}
var cards1org = []Card{{Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}}
var cards2 = []Card{{Suit: "Clubs", Number: 11}, {Suit: "Spades", Number: 3}, {Suit: "Clubs", Number: 13}, {Suit: "Clubs", Number: 14}}
var cards2org = []Card{{Suit: "Clubs", Number: 11}, {Suit: "Spades", Number: 3}, {Suit: "Clubs", Number: 13}, {Suit: "Clubs", Number: 14}}
var cards3 = []Card{{Suit: "Hearts", Number: 3}, {Suit: "Clubs", Number: 13}, {Suit: "Spades", Number: 5}, {Suit: "Diamonds", Number: 10}}
var cards3org = []Card{{Suit: "Hearts", Number: 3}, {Suit: "Clubs", Number: 13}, {Suit: "Spades", Number: 5}, {Suit: "Diamonds", Number: 10}}
var cards4 = []Card{{Suit: "Hearts", Number: 6}}
var cards4org = []Card{{Suit: "Hearts", Number: 6}}

func main() {
	var tests = []struct {
		input    []Card
		inputOrg []Card
		want     bool
	}{
		{cards1, cards1org, false},
		{cards2, cards2org, false},
		{cards3, cards3org, false},
		{cards4, cards4org, false},
	}

	for _, test := range tests {
		if got := Shuffle(test.input); test.want != testEq(test.inputOrg, got) {
			fmt.Println("Check response ", testEq(test.inputOrg, got))
			fmt.Println("Test.want ", test.want)
			fmt.Println("Test input ", test.input)
			fmt.Println("got ", got)
			fmt.Println("orginal Input: ", cards1org)
			fmt.Println("did they shuffle: (%q) = %v", test.input, got)
		}
	}
}

func testEq(a, b []Card) bool { // test to see if the 2 decks are equal

	if a == nil && b == nil {
		fmt.Println("Returning true")
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Shuffle(slc []Card) []Card {
	fmt.Println("Before Shuffle: ", slc)
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
	fmt.Println("After Shuffle: ", slc)

	return slc
}

/*
	var tests = []struct {
		input []Card
		want  bool
	}{
		{cards1, false},
		{cards2, false},
		{cards3, false},
		{cards4, false},
		{cards5, false}, // no matter how you shuffle two of the same, they should be the same
	}

	for _, test := range tests {
		if got := Shuffle(test.input); test.want == testEq(test.input, got) {
			t.Errorf("did they shuffle: (%q) = %v", test.input, got)
		}
	}

*/
