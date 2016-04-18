package wordGame

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"b", true},
		{"dad", true},
		{"daddy", false},
		{"mom", true},
		{"brother", false},
		{"été", true},      // french testing
		{"Rödlöga", false}, // Swedish testing
		{"ö", true},        // more swedish stuff
		{"palindrome", false},
		{"gopher", false},
		{"A man, a plan, a canal; Panama!", true},
		{"A moon, 2 suns, snus noom a", true},
		{"Regelbasisableger", true},
		{"solutomaattimittaamotulos", true},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
