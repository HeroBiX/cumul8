package main

import (
	"testing"
)

var cards1 = []Card{{Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}}
var cards1org = []Card{{Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}}
var cards2 = []Card{{Suit: "Clubs", Number: 11}, {Suit: "Spades", Number: 3}, {Suit: "Clubs", Number: 13}, {Suit: "Clubs", Number: 14}}
var cards2org = []Card{{Suit: "Clubs", Number: 11}, {Suit: "Spades", Number: 3}, {Suit: "Clubs", Number: 13}, {Suit: "Clubs", Number: 14}}
var cards3 = []Card{{Suit: "Hearts", Number: 3}, {Suit: "Clubs", Number: 13}, {Suit: "Spades", Number: 5}, {Suit: "Diamonds", Number: 10}}
var cards3org = []Card{{Suit: "Hearts", Number: 3}, {Suit: "Clubs", Number: 13}, {Suit: "Spades", Number: 5}, {Suit: "Diamonds", Number: 10}}
var cards4 = []Card{{Suit: "Hearts", Number: 6}, {Suit: "Hearts", Number: 6}}
var cards4org = []Card{{Suit: "Hearts", Number: 6}, {Suit: "Hearts", Number: 6}}
var cards4less = []Card{{Suit: "Hearts", Number: 6}}

var tests = []struct {
	input    []Card
	inputOrg []Card
	want     bool
}{
	{cards1, cards1org, false},
	{cards2, cards2org, false},
	{cards3, cards3org, false},
	{cards4, cards4org, true}, // no matter how you do it, shuffeling 2 of the same will be true
	// {cards4, cards4less, false}, // will fail during the second part of the test. Succeed 1st part of the test because the decks are un-even

}

func TestShuffle(t *testing.T) { // test the shuffle function

	for _, test := range tests {
		// 1st part: test to see if the deck was shuffled
		if got := Shuffle(test.input); test.want != testEq(test.inputOrg, got) {
			t.Errorf("They didn't shuffle: (%q) = %v", test.input, got)
		}
		// 2nd part: test to see if any card disappears
		if got := Shuffle(test.input); len(got) != len(test.inputOrg) {
			t.Errorf("Card dissappeard: (%q) = %v", test.input, got)
		}
	}

}

func TestRiffle(t *testing.T) { // test the riffle function

	for _, test := range tests {
		// 1st part: test to see if the deck was shuffled
		if got := Riffle(test.input); test.want != testEq(test.inputOrg, got) {
			t.Errorf("They didn't shuffle: (%q) = %v", test.input, got)
		}
		// 2nd part: test to see if any card disappears
		if got := Riffle(test.input); len(got) != len(test.inputOrg) {
			t.Errorf("Card dissappeard: (%q) = %v", test.input, got)
		}
	}
}

func TestShuffleRiffleShuffle(t *testing.T) { // test the ShuffleRiffleShuffle function

	for _, test := range tests {
		// 1st part: test to see if the deck was shuffled
		if got := ShuffleRiffleShuffle(test.input); test.want != testEq(test.inputOrg, got) {
			t.Errorf("They didn't shuffle: (%q) = %v", test.input, got)
		}
		// 2nd part: test to see if any card disappears
		if got := ShuffleRiffleShuffle(test.input); len(got) != len(test.inputOrg) {
			t.Errorf("Card dissappeard: (%q) = %v", test.input, got)
		}
	}

}

func TestQuarterDeck(t *testing.T) {
	var cards8 = []Card{{Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}}
	var cards52 = []Card{{Suit: "Spades", Number: 2}, {Suit: "Spades", Number: 3}, {Suit: "Spades", Number: 4}, {Suit: "Spades", Number: 5}, {Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}, {Suit: "Spades", Number: 14}, {Suit: "Hearts", Number: 2}, {Suit: "Hearts", Number: 3}, {Suit: "Hearts", Number: 4}, {Suit: "Hearts", Number: 5}, {Suit: "Hearts", Number: 6}, {Suit: "Hearts", Number: 7}, {Suit: "Hearts", Number: 8}, {Suit: "Hearts", Number: 9}, {Suit: "Hearts", Number: 10}, {Suit: "Hearts", Number: 11}, {Suit: "Hearts", Number: 12}, {Suit: "Hearts", Number: 13}, {Suit: "Hearts", Number: 14}, {Suit: "Diamonds", Number: 2}, {Suit: "Diamonds", Number: 3}, {Suit: "Diamonds", Number: 4}, {Suit: "Diamonds", Number: 5}, {Suit: "Diamonds", Number: 6}, {Suit: "Diamonds", Number: 7}, {Suit: "Diamonds", Number: 8}, {Suit: "Diamonds", Number: 9}, {Suit: "Diamonds", Number: 10}, {Suit: "Diamonds", Number: 11}, {Suit: "Diamonds", Number: 12}, {Suit: "Diamonds", Number: 13}, {Suit: "Diamonds", Number: 14}, {Suit: "Clubs", Number: 2}, {Suit: "Clubs", Number: 3}, {Suit: "Clubs", Number: 4}, {Suit: "Clubs", Number: 5}, {Suit: "Clubs", Number: 6}, {Suit: "Clubs", Number: 7}, {Suit: "Clubs", Number: 8}, {Suit: "Clubs", Number: 9}, {Suit: "Clubs", Number: 10}, {Suit: "Clubs", Number: 11}, {Suit: "Clubs", Number: 12}, {Suit: "Clubs", Number: 13}, {Suit: "Clubs", Number: 14}}

	var tests = []struct {
		input1 []Card
		input2 []Card
		want   int
	}{
		{cards8, cards8, 4},
		{cards52, cards52, 26},
	}

	for _, test := range tests {
		if _, _, got := QuarterDeck(len(test.input1)/4, test.input1, test.input2); test.want != len(got) {
			t.Errorf("They didn't take a quarter: (%q, %q) = %v", len(test.input1), len(test.input2), len(got))
		}
	}
}

func TestSplitDeck(t *testing.T) {
	var cards8 = []Card{{Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}}
	var cards52 = []Card{{Suit: "Spades", Number: 2}, {Suit: "Spades", Number: 3}, {Suit: "Spades", Number: 4}, {Suit: "Spades", Number: 5}, {Suit: "Spades", Number: 6}, {Suit: "Spades", Number: 7}, {Suit: "Spades", Number: 8}, {Suit: "Spades", Number: 9}, {Suit: "Spades", Number: 10}, {Suit: "Spades", Number: 11}, {Suit: "Spades", Number: 12}, {Suit: "Spades", Number: 13}, {Suit: "Spades", Number: 14}, {Suit: "Hearts", Number: 2}, {Suit: "Hearts", Number: 3}, {Suit: "Hearts", Number: 4}, {Suit: "Hearts", Number: 5}, {Suit: "Hearts", Number: 6}, {Suit: "Hearts", Number: 7}, {Suit: "Hearts", Number: 8}, {Suit: "Hearts", Number: 9}, {Suit: "Hearts", Number: 10}, {Suit: "Hearts", Number: 11}, {Suit: "Hearts", Number: 12}, {Suit: "Hearts", Number: 13}, {Suit: "Hearts", Number: 14}, {Suit: "Diamonds", Number: 2}, {Suit: "Diamonds", Number: 3}, {Suit: "Diamonds", Number: 4}, {Suit: "Diamonds", Number: 5}, {Suit: "Diamonds", Number: 6}, {Suit: "Diamonds", Number: 7}, {Suit: "Diamonds", Number: 8}, {Suit: "Diamonds", Number: 9}, {Suit: "Diamonds", Number: 10}, {Suit: "Diamonds", Number: 11}, {Suit: "Diamonds", Number: 12}, {Suit: "Diamonds", Number: 13}, {Suit: "Diamonds", Number: 14}, {Suit: "Clubs", Number: 2}, {Suit: "Clubs", Number: 3}, {Suit: "Clubs", Number: 4}, {Suit: "Clubs", Number: 5}, {Suit: "Clubs", Number: 6}, {Suit: "Clubs", Number: 7}, {Suit: "Clubs", Number: 8}, {Suit: "Clubs", Number: 9}, {Suit: "Clubs", Number: 10}, {Suit: "Clubs", Number: 11}, {Suit: "Clubs", Number: 12}, {Suit: "Clubs", Number: 13}, {Suit: "Clubs", Number: 14}}

	var tests = []struct {
		input1 []Card
		want   int
	}{
		{cards8, 4},
		{cards52, 26},
	}

	for _, test := range tests {
		if _, got := SplitDeck(test.input1); test.want != len(got) {
			t.Errorf("They didn't split it: (%q, %q) = %v", len(test.input1), len(got))
		}
	}
}

func testEq(a, b []Card) bool { // test to see if the 2 decks are matching up

	if a == nil && b == nil {
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
