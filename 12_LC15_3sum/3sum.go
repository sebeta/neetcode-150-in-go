// Package threesum solves the "3Sum" problem.
package threesum

import "slices"

// ThreeSum returns all the distinct triplets [nums[i], nums[j], nums[k]] with
// different positions i, j and k whose sum is 0. The result has no repeated
// triplets, and neither the order of the triplets nor the order of the
// numbers inside each one is guaranteed. It sorts nums in place.
//
// It sorts nums and fixes each number nums[i] in turn. Then it looks for a
// pair to the right of i that adds up to -nums[i], using two pointers (one
// moving right and the other moving left, depending on the sum). Repeated
// values of nums[i] and of the left pointer are skipped to avoid duplicate
// triplets.
//
// Time complexity: O(n^2), where n is the length of nums. Sorting takes
// O(n log n), and the two pointers run through nums once for each i.
// Space complexity: O(1) extra space, not counting the result (the sorting
// may use O(log n) of stack space).
//
// NOTE: This is my solution after getting some help from AI and applying the
// same logic as in "Two Sum II" exercise (see ../11_LC167_twosum2).
// It is also one of the optimal solutions from the video.
func ThreeSum(nums []int) [][]int {
	triplets := make([][]int, 0)
	slices.Sort(nums) // Sorting so two pointers can find pairs and duplicates are adjacent
	for i := range nums {
		// Skip repeated values of nums[i] to avoid duplicate triplets:
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		// The pair (l, r) must add up to target so the triplet sums to 0:
		target := -1 * nums[i]
		l := i + 1         // Left
		r := len(nums) - 1 // Right

		for l < r {
			if nums[l]+nums[r] == target {
				triplets = append(triplets, []int{nums[i], nums[l], nums[r]})
				// Changing l and r to keep looking for new triplets:
				l++
				r--
				// Skipping repeated values of l to avoid duplicate triplets:
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				l++
			}
		}
	}

	return triplets
}
