package controllers

import (
	"gopkg.in/mgo.v2"
	"testing"
)

// user Dennis, frej and cumul8 has been added for this test

func TestAddFileName(t *testing.T) {
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	defer s.Close()

	var tests = []struct {
		input    string
		fileName string
		want     error
	}{
		{"cumul8", "bobs.file", nil},
		{"Dennis", "dennis.file", nil},
		{"frej", "secret.txt", nil},
	}

	for _, test := range tests {
		if got := AddFileName(s, test.fileName, test.input); test.want != got {
			t.Errorf("problem adding the file name: (%q) = %v", test.input, got)
		}
	}
}

func TestGetUser(t *testing.T) {
	s, err := mgo.Dial("mongodb://localhost")
	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	defer s.Close()

	var tests = []struct {
		input string
		want  error
	}{
		{"cumul8", nil},
		{"Dennis", nil},
		{"frej", nil},
	}

	for _, test := range tests {
		if _, got := GetUser(test.input, s); test.want != got {
			t.Errorf("Error when getting user info: (%q) = %v", test.input, got)
		}
	}
}
