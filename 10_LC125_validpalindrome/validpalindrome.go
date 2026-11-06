// Package validpalindrome solves the "Valid palindrome" problem.
package validpalindrome

import "unicode"

// IsPalindrome reports whether str is a palindrome, ignoring the case of the
// letters and any character that is not a letter or a digit. A string with
// no letters or digits is considered a palindrome.
//
// It uses two pointers, one at the start and one at the end of str. Each one
// skips the characters that are not alphanumeric, the two characters are
// compared in lowercase, and then the pointers move towards the center.
// It assumes str only contains ASCII characters, so each byte is a character.
//
// Time complexity: O(n), where n is the length of str.
// Space complexity: O(1), since no extra string is created.
//
// NOTE: This is the most optimal solution from the video.
func IsPalindrome(str string) bool {
	l := 0            // Left
	r := len(str) - 1 // Right

	// Stops when the pointers meet: the middle character of an odd-length
	// string would only be compared with itself.
	for l < r {
		// The l < r check in the inner loops keeps the pointers from crossing
		// (or leaving the string) while skipping non-alphanumeric characters.
		for l < r && !unicode.IsLetter(rune(str[l])) && !unicode.IsDigit(rune(str[l])) {
			l++
		}
		for l < r && !unicode.IsLetter(rune(str[r])) && !unicode.IsDigit(rune(str[r])) {
			r--
		}
		if unicode.ToLower(rune(str[l])) != unicode.ToLower(rune(str[r])) {
			return false
		}
		l++
		r--
	}

	return true
}
