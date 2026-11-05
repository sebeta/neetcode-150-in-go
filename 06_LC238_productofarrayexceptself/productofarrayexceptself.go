// Package productofarrayexceptself solves the "Product of array except self" problem.
package productofarrayexceptself

// ProductExceptSelf returns an array where each position i holds the product
// of all the elements of nums except nums[i], without using division.
// It makes two passes over nums: the first stores in each position the product
// of all the elements to its left, and the second multiplies it by the product
// of all the elements to its right, using a single running variable.
//
// Time complexity: O(n), where n is the length of nums.
// Space complexity: O(1) extra space, since the output array does not count.
//
// NOTE: This is my solution after some help from AI.
// This is the same as the optimal one in the video.
func ProductExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	// First, travel from left to right
	prod := 1
	for i := range nums {
		// answer[i] will be multiplied by the prod from previous iteration
		answer[i] = prod
		prod *= nums[i]
	}
	// Then, from right to left
	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		// answer[i] will be multiplied by the prod from previous iteration
		answer[i] *= prod
		prod *= nums[i]
	}

	return answer
}
