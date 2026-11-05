// Package twosum solves the "Two sum" problem.
package twosum

// TwoSum returns the indices of the two numbers in nums that add up to target.
// It looks up each number's complement (target - num) in a map of the numbers
// seen so far, so each element is only visited once. It returns nil if no
// pair adds up to target.
//
// Time complexity: O(n), where n is the length of nums.
// Space complexity: O(n), for the map of seen numbers.
//
//	NOTE: This is the most optimal solution from the video.
func TwoSum(nums []int, target int) []int {
	positionPerNum := make(map[int]int) // Number -> Number's position
	for i, num := range nums {
		complement := target - num
		if compPosition, ok := positionPerNum[complement]; ok {
			return []int{compPosition, i}
		}

		positionPerNum[num] = i
	}

	// nil will never be returned as per problem constraints
	return nil
}
