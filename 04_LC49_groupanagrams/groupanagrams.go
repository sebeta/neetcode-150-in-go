// Package groupanagrams solves the "Group anagrams" problem.
package groupanagrams

// GroupAnagrams groups the strings in strs that are anagrams of each other.
// Each string gets an "anagram sign" (the count of each letter), and strings
// with the same sign end up in the same group. The groups are returned in the
// order in which their first string appears in strs.
// It assumes all strings only contain lowercase English letters (a-z).
//
// Time complexity: O(n * k), where n is the length of strs and k is the
// length of the longest string.
// Space complexity: O(n * k), for the groups and the map of signs.
//
// NOTE: This is my solution. It is very similar to the most optimal solution in the video.
func GroupAnagrams(strs []string) [][]string {
	var anagrams [][]string
	positionPerSign := make(map[[26]int]int)
	for _, str := range strs {
		sign := getAnagramSignature(str)
		if pos, ok := positionPerSign[sign]; ok {
			anagrams[pos] = append(anagrams[pos], str)
			continue
		}

		anagrams = append(anagrams, []string{str})
		positionPerSign[sign] = len(anagrams) - 1
	}

	return anagrams
}

// getAnagramSignature returns the number of times each letter (a-z) appears in str.
// Two strings are anagrams if and only if they have the same sign.
//
// NOTE: The signature could also be a string or a [26]uint8.
func getAnagramSignature(str string) [26]int {
	var sign [26]int
	for _, r := range str {
		sign[r-'a']++
	}

	return sign
}
