// Package longestconsecutive solves the "Longest consecutive sequence" problem.
package longestconsecutive

// LongestConsecutive returns the length of the longest sequence of
// consecutive integers (such as 1, 2, 3, 4) that can be formed with the
// elements of nums, in any order. Duplicates are counted only once.
//
// It puts all the numbers in a set and only counts from the numbers that
// start a sequence (the ones whose previous number, num-1, is not in the
// set). Each number is visited at most twice (once in the outer loop and
// once while counting a sequence), so the nested loops are still linear.
//
// Time complexity: O(n), where n is the length of nums.
// Space complexity: O(n), for the set of numbers.
//
// NOTE: This is my solution (I'm kind of pride of it).
// This is also the most optimal according to the video.
func LongestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	var longestSeq int
	for num := range numSet {
		// Considering the start of a sequence only:
		if _, ok := numSet[num-1]; ok {
			continue
		}
		seqCount := 1
		next := num + 1
		for {
			if _, ok := numSet[next]; !ok {
				if seqCount > longestSeq {
					longestSeq = seqCount
				}
				break
			}
			seqCount++
			next++
		}
	}

	return longestSeq
}
