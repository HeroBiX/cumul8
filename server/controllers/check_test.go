package controllers

import (
	"gopkg.in/mgo.v2"
	"testing"
)

func TestExistingUserName(t *testing.T) {
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	defer s.Close()

	var tests = []struct {
		input      string
		inputLower string
		want       bool
	}{
		{"Bob", "bob", true},
		{"Dennis", "dennis", false},
		{"Frej", "frej", false},
		{"Tor", "tor", true},
		{"Sigyn", "sigyn", true},
	}

	for _, test := range tests {
		if got := ExistingUserName(test.input, test.inputLower, s); test.want != got {
			t.Errorf("Username/Password is too short: (%q) = %v", test.input, got)
		}
	}
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
		if got := EnoughCharacters(test.input, test.number); test.want != got {
			t.Errorf("Username/Password is too short: (%q) = %v", test.input, got)
		}
	}
}

func TestNoFunkyCharacters(t *testing.T) {

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
		if got := NoFunkyCharacters(test.input); test.want != got {
			t.Errorf("Username/Password has funky characters: (%q) = %v", test.input, got)
		}
	}
}
