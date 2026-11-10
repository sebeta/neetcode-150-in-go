// Package longestsubstring solves the "Longest substring without repeating characters" problem.
package longestsubstring

// LengthOfLongestSubstring returns the length of the longest substring of s
// that has no repeated characters. A substring is a continuous part of s, so
// a subsequence such as "pwke" in "pwwkew" doesn't count. It returns 0 for an
// empty string.
//
// It uses a sliding window: start is the position where the current
// substring begins, and the position i being read is where it ends. A map
// keeps the last position of every character seen in s. When a character
// already appeared inside the window (its last position is at or after
// start), the window can't contain both, so start moves to the position right
// after the previous one. Characters that appeared before start are ignored,
// since they are no longer part of the window, so start never moves
// backwards. At each step the window has i-start+1 characters, and the
// biggest one is the answer.
//
// For example, with "abba": at the second "b", start moves from 0 to 2; at the
// last "a", its previous position (0) is before start, so start doesn't move
// and the window "ba" has length 2.
//
// It assumes s only has ASCII characters, so each byte is a character.
//
// Time complexity: O(n), where n is the length of s, since each character is
// visited once.
// Space complexity: O(k), where k is the number of different characters in s
// (at most 95 printable ASCII characters), so it is O(1) in practice.
//
// NOTE: This is my solution, I got some help from AI.
// This is also the most optimal solution from the video.
func LengthOfLongestSubstring(s string) int {
	var longestSubLen int
	// Character -> Character's last position in s
	posPerRune := make(map[rune]int)
	// Position in s where the current substring starts
	var start int
	for i, ch := range s {
		// If ch was already seen inside the current substring, the substring
		// has to start right after that previous position.
		if pos, ok := posPerRune[ch]; ok && pos >= start {
			start = pos + 1
		}
		posPerRune[ch] = i

		longestSubLen = max(longestSubLen, i-start+1)
	}

	return longestSubLen
}
