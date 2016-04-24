package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{1, "1"},
		{3, "fizz"},
		{5, "buzz"},
		{6, "fizz"},
		{7, "7"},
		{15, "fizzbuzz"},
		{26, "26"},
		{27, "fizz"},
		{43, "43"},
		{48, "fizz"},
		{70, "buzz"},
		{78, "fizz"},
		{83, "83"},
		{90, "fizzbuzz"},
		{98, "98"},
		{100, "buzz"},
	}

	for _, test := range tests {
		if got := FizzBuzz(test.input); got != test.want {
			t.Errorf("FizzBuzz(%q) = %v, want %q", test.input, got, test.want)
		}
	}
}
