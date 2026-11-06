// Package twosum2 solves the "Two sum II" problem.
package twosum2

// TwoSum2 returns the 1-indexed positions [index1, index2] (index1 < index2)
// of the two numbers in numbers that add up to target. numbers must be sorted
// in non-decreasing order. It returns nil if there is no such pair.
//
// It uses two pointers, one at the start and one at the end of numbers. If
// the sum of both is bigger than target, the right pointer moves left to get
// a smaller sum; if it is smaller, the left pointer moves right to get a
// bigger one. Every move discards a number that can't be part of the answer.
//
// Time complexity: O(n), where n is the length of numbers.
// Space complexity: O(1), since only two pointers are used.
//
// NOTE: This is my solution after some hints from AI.
// Is is the same as the most optimal in the video.
func TwoSum2(numbers []int, target int) []int {
	l := 0                // Left
	r := len(numbers) - 1 // Right
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}

	// nil will never be returned as per problem constraints
	return nil
}
