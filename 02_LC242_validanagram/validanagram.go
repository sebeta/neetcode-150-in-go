// Package validanagram solves the "Valid anagram" problem.
package validanagram

// ValidAnagram reports whether t is an anagram of s, meaning both strings
// contain the same characters with the same number of repetitions.
// It assumes both strings only contain lowercase English letters (a-z).
//
// Time complexity: O(n), where n is the length of s (and t).
// Space complexity: O(1), since the counter array always has 26 positions.
//
// NOTE: This is the most optimal solution from the video.
func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// There are 26 characters in the English alphabet
	charCounts := make([]int, 26)

	for i := 0; i < len(s); i++ {
		charCounts[s[i]-'a']++
		charCounts[t[i]-'a']--
	}

	for _, c := range charCounts {
		if c != 0 {
			return false
		}
	}

	return true
}

// ValidAnagram2 reports whether t is an anagram of s, meaning both strings
// contain the same characters with the same number of repetitions.
//
// Time complexity: O(n + m), where n and m are the lengths of s and t.
// Space complexity: O(k), where k is the number of distinct characters in s
// and t (at most n + m). It is O(1) if the strings only contain a-z, since
// each map holds at most 26 entries.
//
// NOTE: This is my solution.
func ValidAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sFreqs := getCharacterFrequency(s)
	tFreqs := getCharacterFrequency(t)
	if len(sFreqs) != len(tFreqs) {
		return false
	}

	for ch, sFreq := range sFreqs {
		tFreq := tFreqs[ch]
		if sFreq != tFreq {
			return false
		}
	}

	return true
}

// getCharacterFrequency returns a map with the number of times each
// character appears in s.
func getCharacterFrequency(s string) map[string]int {
	characterFrequency := make(map[string]int)
	for _, r := range s {
		characterFrequency[string(r)]++
	}

	return characterFrequency
}
