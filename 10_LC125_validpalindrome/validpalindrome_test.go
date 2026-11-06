package validpalindrome

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"leetcode example 1", "A man, a plan, a canal: Panama", true},
		{"leetcode example 2", "race a car", false},
		{"leetcode example 3: a single space", " ", true},

		// Small inputs.
		{"single letter", "a", true},
		{"single digit", "7", true},
		{"two equal letters", "aa", true},
		{"two different letters", "ab", false},
		{"odd length palindrome", "aba", true},
		{"even length palindrome", "abba", true},
		{"not a palindrome", "abc", false},
		{"not a palindrome with odd length", "abcda", false},

		// Case is ignored.
		{"same letter in different case", "Aa", true},
		{"mixed case palindrome", "Racecar", true},
		{"all uppercase palindrome", "RACECAR", true},
		{"mixed case that is not a palindrome", "Race", false},

		// Strings that are empty after removing non-alphanumeric characters.
		{"only spaces", "   ", true},
		{"only punctuation", ".,!?", true},
		{"only symbols", "@#$%^&*()", true},
		{"single punctuation", ",", true},

		// Digits.
		{"digits palindrome", "12321", true},
		{"digits that are not a palindrome", "12345", false},
		{"digit and letter that look alike after lowercase", "0P", false},
		{"letters and digits palindrome", "A1b2b1a", true},
		{"letters and digits that are not a palindrome", "1a2", false},

		// Non-alphanumeric characters are removed.
		{"punctuation at the end", "a.", true},
		{"punctuation at the start", ",a", true},
		{"punctuation at both ends", "!a!", true},
		{"punctuation in the middle", "ab.ba", true},
		{"spaces between letters", "a b a", true},
		{"underscore is not alphanumeric", "ab_a", true},
		{"punctuation hides a mismatch", "a,b", false},
		{"punctuation around a mismatch", ".a.b.", false},
		{"spaces around a mismatch", " a b ", false},

		// Characters next to letters and digits in the ASCII table. They are
		// not alphanumeric, so they must be removed.
		{"@ is the character before A", "a@", true},
		{"[ is the character after Z", "a[", true},
		{"` is the character before a", "a`", true},
		{"{ is the character after z", "a{", true},
		{"/ is the character before 0", "0/", true},
		{": is the character after 9", "0:", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.s); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// TestIsPalindromeLargeInput uses inputs close to the maximum allowed length
// (2 * 10^5 characters). A solution that is not linear (for example, one that
// removes characters from the string one by one) takes much longer here.
func TestIsPalindromeLargeInput(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			"long palindrome of one letter",
			strings.Repeat("a", 200_000),
			true,
		},
		{
			"long palindrome with punctuation and spaces",
			strings.Repeat("A man, a plan, a canal: Panama ", 5_000), // 155,000 characters
			true,
		},
		{
			"long string with the mismatch in the middle",
			strings.Repeat("a", 100_000) + "b" + strings.Repeat("a", 99_999),
			false,
		},
		{
			"long string with the mismatch at the ends",
			"b" + strings.Repeat("a", 199_998) + "c",
			false,
		},
		{
			"long string of only punctuation",
			strings.Repeat(".", 200_000),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.s); got != tt.want {
				t.Errorf("IsPalindrome(large input) = %v, want %v", got, tt.want)
			}
		})
	}
}
