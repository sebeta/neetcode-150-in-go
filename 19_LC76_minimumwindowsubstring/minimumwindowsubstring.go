// Package minimumwindowsubstring solves the "Minimum window substring" problem.
package minimumwindowsubstring

// MinWindow returns the shortest substring of s that includes every character
// of t, counting the duplicates (if t has two 'a', the substring needs at
// least two 'a'). It can have other characters too. It returns the empty
// string if there is no such substring. Both strings only have uppercase and
// lowercase English letters, and the answer is assumed to be unique.
//
// It uses a sliding window s[l:r] that grows and shrinks. It starts with the
// first len(t) letters of s, because a shorter window can't include all the
// letters of t. A map has how many times each letter is needed by t, and
// another one has how many times each letter is inside the window. While the
// window doesn't have everything t needs, the next letter of s enters from the
// right (r moves). As soon as it does, its length is compared with the best
// one found, and the letter on the left leaves (l moves) to look for a shorter
// window. When no letters are left to enter and the window is not valid
// anymore, the search ends. Only the start and the length of the best window
// are kept, and the substring is cut once at the end.
//
// For example, with s="ADOBECODEBANC" and t="ABC": the first valid window is
// "ADOBEC"; later the window "BANC" is found, which is shorter, and no shorter
// valid window exists, so it returns "BANC".
//
// Time complexity: O(m + n), where m is the length of s and n the length of t.
// Counting t takes O(n). Then every step moves l or r, so there are at most
// 2m steps, and each step checks at most 52 different letters (a constant).
// Space complexity: O(1), since the maps never have more than 52 letters.
//
// NOTE: This is my solution after a lot of help from AI.
// I know there's room for improvement, but this is the solution I can
// understand faster reading the code.
func MinWindow(s, t string) string {
	if len(t) > len(s) {
		return ""
	}

	tSign := make(map[byte]int)
	windowSign := make(map[byte]int)
	for i := range t {
		tSign[t[i]]++
		windowSign[s[i]]++ // Initial value for window sign
	}

	// The window is s[l:r]. Only the start and the length of the best window
	// are kept, so s is cut once at the end instead of copying every window.
	var bestStart, bestLen int
	l, r := 0, len(t) // Left and right
	for r <= len(s) {
		if isValidWindow(tSign, windowSign) {
			if bestLen == 0 || r-l < bestLen {
				bestStart, bestLen = l, r-l
			}
			windowSign[s[l]]-- // s[l] is the letter that leaves the window
			l++
		} else {
			if r < len(s) {
				windowSign[s[r]]++ // s[r] is the letter that enters the window
			}
			r++ // When r == len(s) there are no letters left, so the loop ends
		}
	}

	return s[bestStart : bestStart+bestLen]
}

// isValidWindow reports whether the window has, for every letter that is
// needed, at least as many times as need says. Letters of the window that are
// not needed, and extra copies of the needed ones, don't matter.
func isValidWindow(need, have map[byte]int) bool {
	for ch, n := range need {
		if have[ch] < n {
			return false
		}
	}

	return true
}
