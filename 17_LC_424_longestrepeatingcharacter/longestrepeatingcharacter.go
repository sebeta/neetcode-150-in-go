// Package longestrepeatingcharacter solves the "Longest repeating character replacement" problem.
package longestrepeatingcharacter

// CharacterReplacement returns the length of the longest substring of s that
// can be turned into a run of one repeated letter by replacing at most k
// characters.
//
// It uses a sliding window that goes from start to the position i being read,
// and a map with how many times each letter appears inside the window. A
// window is valid when its length minus the count of its most frequent letter
// is at most k: those are the characters that would have to be replaced to
// make the whole window the same letter. In each step the letter at i enters
// the window and the highest count is calculated again. If the window is
// valid, its length is compared with the best one found. If it is not, the
// letter at start leaves the window and start moves one step, so the window
// keeps its size and slides to the right.
//
// For example, with "ABAB" and k=1: the windows "A", "AB" and "ABA" are valid
// (for "ABA", 3-2=1), so the best length is 3. When the last "B" enters,
// "ABAB" needs 4-2=2 replacements, which is more than k, so the first "A"
// leaves and the window becomes "BAB". The answer is 3.
//
// Time complexity: O(n), where n is the length of s. Finding the highest
// count looks at the map in every step, but it has at most 26 letters.
// Space complexity: O(1), since the map never has more than 26 letters.
//
// NOTE: This exercise blew my mind. This is my solution after a lot of help
// from AI. I should come back to this one in the future.
func CharacterReplacement(s string, k int) int {
	characterFreq := make(map[rune]int)
	var longestWindow int
	var start int
	for i, ch := range s {
		characterFreq[ch]++
		var maxFreq int
		for _, freq := range characterFreq {
			maxFreq = max(maxFreq, freq)
		}
		windowLen := i + 1 - start
		isReplaceable := (windowLen - maxFreq) <= k
		if isReplaceable {
			longestWindow = max(longestWindow, windowLen)
		} else {
			characterFreq[rune(s[start])]--
			start++ // Move window
		}

	}
	return longestWindow
}

// CharacterReplacement2 returns the length of the longest substring of s that
// can be turned into a run of one repeated letter by replacing at most k
// characters. It is the same idea as CharacterReplacement with these
// improvements:
//   - The letter counts are in a [26]int array indexed by letter, instead of
//     a map, so there is no hashing and nothing is allocated.
//   - It doesn't look for the highest frequency in every step: maxFreq only
//     grows when a letter enters the window and is never lowered. A window
//     can only be better than the best one found if maxFreq grows, so a
//     stale maxFreq never gives a wrong answer; it just keeps the window from
//     shrinking.
//   - The window doesn't need to be checked more than once per step: it
//     slides to the right (one letter enters and, if needed, one leaves) and
//     its size never decreases, so the final size is the answer.
//
// A window is valid when its length minus the count of its most frequent
// letter is at most k, because those are the characters to replace.
//
// Time complexity: O(n), where n is the length of s.
// Space complexity: O(1), since the array has 26 positions.
//
// It assumes s only has uppercase English letters.
//
// NOTE: This is the most optimal solution according to the AI.
func CharacterReplacement2(s string, k int) int {
	var count [26]int
	var maxFreq, start int
	for i := 0; i < len(s); i++ {
		count[s[i]-'A']++
		maxFreq = max(maxFreq, count[s[i]-'A'])

		// If the window needs more than k replacements, slide it: the letter
		// on the left leaves and the window keeps its size.
		if (i-start+1)-maxFreq > k {
			count[s[start]-'A']--
			start++
		}
	}

	// The window never shrinks, so its size at the end is the best one.
	return len(s) - start
}
