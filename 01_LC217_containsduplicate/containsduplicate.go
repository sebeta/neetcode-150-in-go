// Package containsduplicate solves the "Contains duplicate" problem.
package containsduplicate

// ContainsDuplicate reports whether any value appears at least twice in nums.
//
// Time complexity: O(n), where n is the length of nums.
// Space complexity: O(n), for the set of seen numbers.
//
// NOTE: This is my solution. It is very similar to the most optimal solution in the video.
func ContainsDuplicate(nums []int) bool {
	seenNumbers := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := seenNumbers[num]; ok {
			return true
		}
		seenNumbers[num] = struct{}{}
	}

	return false
}
