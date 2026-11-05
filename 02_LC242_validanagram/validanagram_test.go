package validanagram

import "testing"

type anagramCase struct {
	name string
	s, t string
	want bool
}

// lowercaseCases only use characters a-z, so every implementation must pass them.
var lowercaseCases = []anagramCase{
	{"leetcode example 1", "anagram", "nagaram", true},
	{"leetcode example 2", "rat", "car", false},
	{"single char equal", "a", "a", true},
	{"single char different", "a", "b", false},
	{"identical strings", "listen", "listen", true},
	{"simple anagram", "listen", "silent", true},
	{"different lengths, t longer", "ab", "abc", false},
	{"different lengths, s longer", "abc", "ab", false},
	{"same letters, different counts", "aab", "abb", false},
	{"same length, different letters", "abcd", "abce", false},
	{"repeated letters anagram", "aabbcc", "abcabc", true},
	{"reversed string", "abcdef", "fedcba", true},
	{"all same letter", "aaaa", "aaaa", true},
	{"all same letter, different count", "aaaa", "aaa", false},
	{"long anagram", "abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", true},
}

// unicodeCases cover the follow up of the problem (inputs with Unicode
// characters). They are only valid for implementations that don't assume
// a-z (ValidAnagram2).
var unicodeCases = []anagramCase{
	{"unicode anagram", "añb", "bña", true},
	{"unicode mismatch", "añb", "anb", false},
}

func runCases(t *testing.T, fnName string, fn func(s, t string) bool, cases []anagramCase) {
	t.Helper()
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn(tt.s, tt.t); got != tt.want {
				t.Errorf("%s(%q, %q) = %v, want %v", fnName, tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func TestValidAnagram(t *testing.T) {
	runCases(t, "ValidAnagram", ValidAnagram, lowercaseCases)
}

func TestValidAnagram2(t *testing.T) {
	runCases(t, "ValidAnagram2", ValidAnagram2, lowercaseCases)
	runCases(t, "ValidAnagram2", ValidAnagram2, unicodeCases)
}
