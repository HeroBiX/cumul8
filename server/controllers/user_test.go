package controllers

import (
	"gopkg.in/mgo.v2"
	"testing"
)

func TestListFiles(t *testing.T) {
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
		if got := ListFiles(test.input, test.inputLower, s); test.want != got {
			t.Errorf("Username/Password is too short: (%q) = %v", test.input, got)
		}
	}
}

func TestAddFileName(t *testing.T) {
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
		if got := AddFileName(test.input, test.number); test.want != got {
			t.Errorf("Username/Password is too short: (%q) = %v", test.input, got)
		}
	}
}

func TestGetUser(t *testing.T) {
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
		if got := GetUser(test.input, test.number); test.want != got {
			t.Errorf("Username/Password is too short: (%q) = %v", test.input, got)
		}
	}
}
