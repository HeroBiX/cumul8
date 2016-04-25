package main

import (
	"testing"
)

func TestCheckCreatingUsers(t *testing.T) {
	// Figure out how to connect to the database and check if usernames are being used or not

}

func TestEnoughCharacters(t *testing.T) {
	var tests = []struct {
		input  string
		number int
		want   bool
	}{
		{"Bob", 1, true},
		{"Dennis", 3, true},
		{"Frej", 5, false},
		{"Tor", 2, true},
		{"Sigyn", 2, true},
	}

	for _, test := range tests {
		if got := Stest.input; test.want != got {
			t.Errorf("Username/Password is too short: (%q) = %v", test.input, got)
		}
	}

}

func TestNoFunkyCharacters(t *testing.T) { // test the riffle function

	var tests = []struct {
		input string
		want  bool
	}{
		{"Bob", true},
		{"#$!4", false},
		{"Frej", true},
		{"Tor!", false},
		{"_Sigyn_", false},
		{"The Moose", false},
	}

	for _, test := range tests {
		if got := Stest.input; test.want != got {
			t.Errorf("Username/Password has funky characters: (%q) = %v", test.input, got)
		}
	}

}

func TestShuffleRiffleShuffle(t *testing.T) { // test the ShuffleRiffleShuffle function

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
