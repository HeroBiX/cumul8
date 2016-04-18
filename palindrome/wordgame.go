package wordGame

import "unicode"

// palindrome check if a words is writen the same forward and backwards
// program is ignoring case and non-letters as #@$ 231
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) { // checks if r is a letter
			letters = append(letters, unicode.ToLower(r)) // if it is a letter, make it lowercase and add it to letters
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] { // check if first and last letter is the same and then walking inwords checking if the letters line up
			return false // if one of the letters doesn't match, return false
		}
	}
	return true // no error, return True, the world is Palindrome! eat a cookie! =)
}
