// Package permutationinstring solves the "Permutation in string" problem.
package permutationinstring

// CheckInclusion reports whether s2 contains a permutation of s1, which means
// that some substring of s2 has the same letters as s1, the same number of
// times each, in any order. Both strings only have lowercase English letters.
//
// It uses a sliding window of the same length as s1 over s2, and compares the
// "sign" of s1 and of the window: an array with the number of times each
// letter appears. Two strings are permutations of each other if and only if
// their signs are equal. The sign of the first window is calculated once; then
// the window slides one position at a time, so the letter that leaves has its
// count lowered and the letter that enters has its count raised, and the signs
// are compared again. If s1 is longer than s2, no window fits and it returns
// false.
//
// For example, with s1="ab" and s2="eidbaooo": the signs of the windows "ei",
// "id", "db" and so on are different from the sign of "ab", but the window "ba"
// has one 'a' and one 'b', so it returns true.
//
// Time complexity: O(n), where n is the length of s2. Each slide changes two
// counts, and comparing the two arrays costs 26 steps, which is constant.
// Space complexity: O(1), since the signs are arrays of 26 positions.
//
// NOTE: This is my solution. I used a little help from AI.
// I'm pretty sure this is the most optimal solution.
func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Sign := [26]int{}
	windowSign := [26]int{}
	for i := range s1 {
		s1Sign[s1[i]-'a']++
		windowSign[s2[i]-'a']++ // The first window of s2
	}
	if s1Sign == windowSign {
		return true
	}

	// Slide the window: s2[j-len(s1)] leaves and s2[j] enters.
	for j := len(s1); j < len(s2); j++ {
		windowSign[s2[j-len(s1)]-'a']--
		windowSign[s2[j]-'a']++
		if s1Sign == windowSign {
			return true
		}
	}

	return false
}
